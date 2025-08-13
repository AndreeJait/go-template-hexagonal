package auth

import (
	"context"
	"github.com/AndreeJait/go-template-hexagonal/internal/domain"
)

type UserRepo interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type UseCase interface {
	GetUserByEmail(ctx context.Context, req GetUserByEmailRequest) (*domain.User, error)
}
