# go-template-hexagonal

Microservice **Template** untuk template — dibangun dengan **Go (Echo v4)**, **PostgreSQL (pgxpool)**, **sqlc** (type-safe queries), dan **sql-migrate** (CLI migrasi terpadu di `cmd/migration`). Arsitektur mengikuti **Hexagonal (Ports & Adapters)**.

## ✨ Fitur
- 🧱 Hexagonal: domain/usecase terpisah dari HTTP/DB
- 🌐 Echo v4 HTTP server
- 🐘 PostgreSQL dengan pgxpool
- 🧬 `sqlc` untuk generate kode query yang type-safe
- 🧭 Migrasi DB via `cmd/migration` (`new | up | down | redo | status`) baca **YAML** env config
- ⚙️ Config per environment di `files/config/<env>/config.yaml`

## 📂 Struktur Proyek (ringkas)
```
.
├─ cmd/
│  ├─ migration/          # CLI migrasi (new/up/down/redo/status)
│  └─ service/            # Entrypoint HTTP (Echo)
├─ files/
│  ├─ config/
│  │  ├─ development/config.yaml
│  │  ├─ production/config.yaml
│  │  └─ staging/config.yaml
│  ├─ migrations/         # *.up.sql / *.down.sql
│  └─ sqlc.yaml           # konfigurasi sqlc
├─ internal/
│  ├─ adapters/
│  │  ├─ inbound/http/    # handler & routes (auth, dll)
│  │  └─ outbound/db/
│  │     └─ postgres/
│  │        └─ sqlc/      # hasil generate sqlc
│  ├─ constant/           # konstanta umum (APP_MODE, error codes, dll)
│  ├─ domain/             # entities
│  ├─ infrastructure/     # config loader, postgres, DI
│  └─ usecase/            # services + ports (auth, dll)
├─ queries/               # file SQL untuk sqlc
├─ Makefile               # run, build, sqlc, migrations
├─ go.mod
└─ go.sum
```

## ⚙️ Konfigurasi
Edit `files/config/<env>/config.yaml` (contoh dev):
```yaml
services:
  host: localhost
  port: 8080
  name: "go-template-svc"

db:
  host: localhost
  port: 5432
  name: template
  user: postgres
  password: template

migrations:
  dir: files/migrations
```
Pilih env dengan `APP_ENV` (`development` default).  
Override `DB_PASSWORD` via env var bila loader mendukung.

## 🏗️ Makefile Commands
Lihat file `Makefile` untuk list lengkap.

## 🚀 Jalankan
1) Generate sqlc
```bash
make sqlc
```
2) Buat migration baru
```bash
make migrate-new
```
3) Apply migrations
```bash
make migrate-up
```
4) Build & run service
```bash
make run
```

## 🧬 Lisensi
[MIT](license)
