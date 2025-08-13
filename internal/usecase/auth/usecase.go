package auth

import (
	"context"
	"github.com/AndreeJait/go-template-hexagonal/internal/domain"
)

type useCase struct {
	userRepo UserRepo
}

func NewUseCase(userRepo UserRepo) UseCase {
	return &useCase{userRepo: userRepo}
}

func (u *useCase) GetUserByEmail(ctx context.Context, req GetUserByEmailRequest) (*domain.User, error) {
	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
