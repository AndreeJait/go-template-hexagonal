package db

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db/postgres/permission"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db/postgres/role"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db/postgres/user"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/db"
)

type Repository struct {
	UserRepo       *user.Repository
	RoleRepo       *role.Repository
	PermissionRepo *permission.Repository
	TxManager      *db.TxManager
}
