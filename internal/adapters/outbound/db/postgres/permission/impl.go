package permission

import (
	"context"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db/postgres/sqlc"
	"github.com/AndreeJait/go-template-hexagonal/internal/domain"
	"github.com/AndreeJait/go-template-hexagonal/internal/utils"

	"github.com/AndreeJait/go-utility/tracer"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	q *sqlc.Queries
}

func (r Repository) GetAllPermissionsByRoleID(ctx context.Context, roleID int64) ([]domain.PermissionSimplified, error) {
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(r.GetAllPermissionsByRoleID))
	defer span.End()
	var permissions []domain.PermissionSimplified
	permissionDb, err := r.q.GetAllPermissionsByRoleID(ctx, roleID)
	if err != nil {
		return permissions, err
	}
	err = utils.ObjectToObject(&permissionDb, &permissions)
	return permissions, err
}

func NewPermissionRepository(q *pgxpool.Pool) *Repository {
	return &Repository{q: sqlc.New(q)}
}
