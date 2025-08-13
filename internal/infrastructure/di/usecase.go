package di

import "github.com/AndreeJait/go-template-hexagonal/internal/usecase/auth"

func (s wire) initUseCase() {
	s.uc.AuthUc = auth.NewUseCase(s.repo.UserRepo)
}
