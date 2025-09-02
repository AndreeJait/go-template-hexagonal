package user

import (
	"context"
	"github.com/AndreeJait/go-template-hexagonal/internal/domain"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/config"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/db"

	"github.com/AndreeJait/go-utility/loggerw"
	"github.com/AndreeJait/go-utility/tracer"
)

type useCase struct {
	log       loggerw.Logger
	cfg       *config.Config
	repo      Repo
	txManager *db.TxManager
}

func (u useCase) GetUserById(ctx context.Context, id int64) (domain.UserSimplified, error) {
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(u.GetUserById))
	defer span.End()
	return u.repo.GetUserById(ctx, id)
}

func NewUseCase(cfg *config.Config, log loggerw.Logger, repo Repo, txManager *db.TxManager) UseCase {
	return &useCase{
		cfg:       cfg,
		repo:      repo,
		txManager: txManager,
		log:       log,
	}
}
