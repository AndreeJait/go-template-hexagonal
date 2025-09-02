package di

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/email"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/config"
	"github.com/AndreeJait/go-template-hexagonal/internal/usecase"
	"github.com/AndreeJait/go-utility/loggerw"
	"github.com/jackc/pgx/v5/pgxpool"
)

type wire struct {
	pg  *pgxpool.Pool
	cfg *config.Config
	log loggerw.Logger

	uc   *usecase.UseCase
	repo *db.Repository

	email email.Email
}

func NewWire(cfg *config.Config, pg *pgxpool.Pool, log loggerw.Logger, e email.Email) wire {
	return wire{
		cfg:   cfg,
		pg:    pg,
		log:   log,
		repo:  &db.Repository{},
		uc:    &usecase.UseCase{},
		email: e,
	}
}
