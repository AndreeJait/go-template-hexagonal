package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"db"`
	Migrations struct {
		Dir string `yaml:"dir"`
	} `yaml:"migrations"`
}

func loadConfig(env string) (*Config, error) {
	if env == "" {
		if v := os.Getenv("APP_ENV"); v != "" {
			env = v
		} else {
			env = "development"
		}
	}
	path := fmt.Sprintf("files/config/%s/config.yaml", env)
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open config: %w", err)
	}
	defer f.Close()

	var c Config
	if err := yaml.NewDecoder(f).Decode(&c); err != nil {
		return nil, fmt.Errorf("parse yaml: %w", err)
	}
	if c.Migrations.Dir == "" {
		c.Migrations.Dir = "files/migrations"
	}
	return &c, nil
}

func main() {
	cmd := flag.String("cmd", "status", "migration command: up | down | redo | status | new")
	steps := flag.Int("n", 0, "number of steps (for down/redo); 0 means all for up")
	env := flag.String("env", "", "APP_ENV override (development|production)")
	name := flag.String("name", "", "migration name (for cmd=new)")
	flag.Parse()

	cfg, err := loadConfig(*env)
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("open db: ", err)
	}
	defer db.Close()

	src := &migrate.FileMigrationSource{Dir: cfg.Migrations.Dir}
	migrate.SetTable("gorp_migrations")

	switch strings.ToLower(*cmd) {
	case "up":
		if *steps <= 0 {
			n, err := migrate.Exec(db, "postgres", src, migrate.Up)
			if err != nil {
				log.Fatal("migrate up:", err)
			}
			log.Printf("Applied %d migration(s)\n", n)
		} else {
			n, err := migrate.ExecMax(db, "postgres", src, migrate.Up, *steps)
			if err != nil {
				log.Fatal("migrate up:", err)
			}
			log.Printf("Applied %d migration(s)\n", n)
		}

	case "down":
		n := *steps
		if n <= 0 {
			n = 1
		}
		n, err := migrate.ExecMax(db, "postgres", src, migrate.Down, n)
		if err != nil {
			log.Fatal("migrate down:", err)
		}
		log.Printf("Rolled back %d migration(s)\n", n)

	case "redo":
		n := *steps
		if n <= 0 {
			n = 1
		}
		for i := 0; i < n; i++ {
			if _, err = migrate.Exec(db, "postgres", src, migrate.Down); err != nil {
				log.Fatal("redo down:", err)
			}
			if _, err = migrate.Exec(db, "postgres", src, migrate.Up); err != nil {
				log.Fatal("redo up:", err)
			}
		}
		log.Printf("Redone %d migration(s)\n", n)

	case "status":
		records, err := migrate.GetMigrationRecords(db, "postgres")
		if err != nil {
			log.Fatal("status:", err)
		}
		applied := map[string]bool{}
		for _, r := range records {
			applied[r.Id] = true
		}
		migs, err := src.FindMigrations()
		if err != nil {
			log.Fatal("find:", err)
		}
		log.Println("Migration status:")
		for _, m := range migs {
			state := "PENDING"
			if applied[m.Id] {
				state = "APPLIED"
			}
			log.Printf(" - %s : %s\n", m.Id, state)
		}

	case "new":
		if *name == "" {
			log.Fatal("migration name is required for cmd=new")
		}
		if err := os.MkdirAll(cfg.Migrations.Dir, 0755); err != nil {
			log.Fatalf("create migrations dir: %v", err)
		}
		timestamp := time.Now().Format("20060102150405")
		base := fmt.Sprintf("%s_%s", timestamp, *name)
		upPath := filepath.Join(cfg.Migrations.Dir, base+".up.sql")
		downPath := filepath.Join(cfg.Migrations.Dir, base+".down.sql")
		if err := os.WriteFile(upPath, []byte("-- +migrate Up\n"), 0644); err != nil {
			log.Fatal("write up file:", err)
		}
		if err := os.WriteFile(downPath, []byte("-- +migrate Down\n"), 0644); err != nil {
			log.Fatal("write down file:", err)
		}
		log.Printf("Created %s\n", upPath)
		log.Printf("Created %s\n", downPath)
	case "fresh":
		log.Println("Dropping all tables...")

		// 1. Drop schema (Postgres-specific)
		_, err := db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
		if err != nil {
			log.Fatalf("failed to reset schema: %v", err)
		}
		log.Println("Schema dropped and recreated.")

		// 2. Re-run all migrations up
		n, err := migrate.Exec(db, "postgres", src, migrate.Up)
		if err != nil {
			log.Fatalf("failed to re-apply migrations: %v", err)
		}
		log.Printf("Applied %d fresh migrations\n", n)

	default:
		log.Fatalf("unknown cmd: %s (use up|down|redo|status|new)", *cmd)
	}
}
