package usecase

import "github.com/AndreeJait/go-template-hexagonal/internal/usecase/auth"

type UseCase struct {
	AuthUc auth.UseCase
}
