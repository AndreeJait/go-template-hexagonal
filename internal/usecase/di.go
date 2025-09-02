package usecase

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/usecase/auth"
	"github.com/AndreeJait/go-template-hexagonal/internal/usecase/user"
)

type UseCase struct {
	AuthUc auth.UseCase
	UserUc user.UseCase
}
