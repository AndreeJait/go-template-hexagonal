package di

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db/postgres/permission"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db/postgres/role"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db/postgres/user"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/db"
)

func (s wire) initRepository() {
	s.repo.UserRepo = user.NewUserRepository(s.pg)
	s.repo.RoleRepo = role.NewRoleRepository(s.pg)
	s.repo.PermissionRepo = permission.NewPermissionRepository(s.pg)

	s.repo.TxManager = db.NewTxManager(s.pg)
}
