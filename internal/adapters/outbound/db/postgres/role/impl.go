package role

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

func (r Repository) GetRoles(ctx context.Context, offset, limit int32) (roles []domain.RoleSimplified, err error) {
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(r.GetRoleByID))
	defer span.End()
	rolesDb, err := r.q.GetRoles(ctx, sqlc.GetRolesParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return roles, err
	}
	err = utils.ObjectToObject(&rolesDb, &roles)
	return roles, err
}

func (r Repository) GetRoleByID(ctx context.Context, roleId int64) (resp domain.RoleSimplified, err error) {
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(r.GetRoleByID))
	defer span.End()
	var role domain.RoleSimplified
	roleDb, err := r.q.GetRoleById(ctx, roleId)
	if err != nil {
		return role, err
	}
	err = utils.ObjectToObject(&roleDb, &role)
	return role, err
}

func NewRoleRepository(q *pgxpool.Pool) *Repository {
	return &Repository{q: sqlc.New(q)}
}
