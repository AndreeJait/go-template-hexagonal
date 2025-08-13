package db

import (
	"context"
	"fmt"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func NewPostgres(ctx context.Context, cfg config.DB) (*pgxpool.Pool, error) {

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	conf, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}
	conf.MaxConns = 20
	conf.MinConns = 2
	conf.MaxConnLifetime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, err
	}

	ctxPing, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := pool.Ping(ctxPing); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}
