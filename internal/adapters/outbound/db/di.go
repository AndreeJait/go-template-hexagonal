package db

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db/postgres/permission"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db/postgres/role"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db/postgres/user"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/db"
)

type Repository struct {
	// UserRepo
	UserRepo *user.Repository
	// TxManager
	TxManager *db.TxManager
	// RoleRepo
	RoleRepo *role.Repository
	// PermissionRepo
	PermissionRepo *permission.Repository
}
