package db

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db/postgres/user"
)

type Repository struct {
	UserRepo *user.Repository
}
