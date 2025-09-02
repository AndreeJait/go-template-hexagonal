SHELL := /usr/bin/env bash
APP_ENV ?= development
BIN_NAME := service
BIN_DIR := bin

.PHONY: run build sqlc swag migrate-new migrate-up migrate-down migrate-redo migrate-status

build:
	@echo "Building binary..."
	@mkdir -p $(BIN_DIR)
	APP_ENV=$(APP_ENV) go build -o ./$(BIN_DIR)/$(BIN_NAME) ./cmd/service

run: build
	@echo "Running service from repo root..."
	@APP_ENV=$(APP_ENV) ./$(BIN_DIR)/$(BIN_NAME)

# Ensure sqlc is installed
ensure-sqlc:
	@command -v sqlc >/dev/null 2>&1 || { \
		echo "Installing sqlc..."; \
		CGO_ENABLED=0 go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest; \
	}

# Run sqlc generate
sqlc: ensure-sqlc
	@echo "Generating SQL code..."
	sqlc generate -f files/sqlc.yaml


# Ensure swag is installed (as before)
ensure-swag:
	@command -v swag >/dev/null 2>&1 || { \
		echo "Installing swag..."; \
		CGO_ENABLED=0 go install github.com/swaggo/swag/cmd/swag@latest; \
	}

# Generate Swagger docs
swag: ensure-swag
	@echo "Generating Swagger docs..."
	swag init \
		-g cmd/service/main.go \
		-d . \
		--parseDependency --parseInternal \
		--exclude ./bin \
		--exclude ./internal/infrastructure/swagger/docs \
		--exclude ./internal/adapters/outbound/postgres/sqlc \
		-o internal/infrastructure/swagger/docs

.PHONY: migrate-new
migrate-new:
	@read -p "Enter migration name: " name; \
	[ -n "$$name" ] || { echo "Migration name cannot be empty"; exit 1; }; \
	APP_ENV=$(APP_ENV) go run ./cmd/migration -cmd=new -name=$$name

.PHONY: migrate-up
migrate-up:
	APP_ENV=$(APP_ENV) go run ./cmd/migration -cmd=up

.PHONY: migrate-fresh
migrate-fresh:
	APP_ENV=$(APP_ENV) go run ./cmd/migration -cmd=fresh

.PHONY: migrate-down
migrate-down:
	APP_ENV=$(APP_ENV) go run ./cmd/migration -cmd=down

.PHONY: migrate-redo
migrate-redo:
	APP_ENV=$(APP_ENV) go run ./cmd/migration -cmd=redo

.PHONY: migrate-status
migrate-status:
	APP_ENV=$(APP_ENV) go run ./cmd/migration -cmd=status