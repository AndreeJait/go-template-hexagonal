package di

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/usecase/auth"
	"github.com/AndreeJait/go-template-hexagonal/internal/usecase/user"
)

func (s wire) initUseCase() {
	s.uc.AuthUc = auth.NewUseCase(s.cfg,
		s.repo.UserRepo,
		s.repo.RoleRepo,
		s.repo.PermissionRepo,
		s.repo.TxManager,
		s.email,
		s.log)

	s.uc.UserUc = user.NewUseCase(s.cfg,
		s.log, s.repo.UserRepo, s.repo.TxManager)
}
