package user

import (
	"context"
	"github.com/AndreeJait/go-template-hexagonal/internal/domain"
)

type Repo interface {
	GetUserById(ctx context.Context, id int64) (domain.UserSimplified, error)
}

type UseCase interface {
	GetUserById(ctx context.Context, id int64) (domain.UserSimplified, error)
}
