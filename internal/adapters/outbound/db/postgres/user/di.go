package user

import (
	"context"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db/postgres/sqlc"
	"github.com/AndreeJait/go-template-hexagonal/internal/constant"
	"github.com/AndreeJait/go-template-hexagonal/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Repository struct {
	q *sqlc.Queries
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	u, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, constant.ErrUserNotFound
		}
		return nil, err
	}
	return &domain.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
		CreateAt: u.CreatedAt.Time,
		UpdateAt: u.UpdatedAt.Time,
	}, nil
}

func NewUserRepository(q *pgxpool.Pool) *Repository {
	return &Repository{q: sqlc.New(q)}
}
