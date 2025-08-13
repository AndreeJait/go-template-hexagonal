package di

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db/postgres/user"
)

func (s wire) initRepository() {
	s.repo.UserRepo = user.NewUserRepository(s.pg)
}
