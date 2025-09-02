package user

import (
	"context"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/db/postgres/sqlc"
	"github.com/AndreeJait/go-template-hexagonal/internal/constant"
	"github.com/AndreeJait/go-template-hexagonal/internal/domain"
	"github.com/AndreeJait/go-template-hexagonal/internal/utils"

	"github.com/AndreeJait/go-utility/tracer"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Repository struct {
	q *sqlc.Queries
}

func (r *Repository) UpdateUserPasswordPinTokenActivationStatusByUserIDParams(ctx context.Context, user domain.User, tx pgx.Tx) error {
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(r.UpdateUserPasswordPinTokenActivationStatusByUserIDParams))
	defer span.End()

	q := r.q
	if tx != nil {
		q = sqlc.New(tx)
	}
	updateParam := sqlc.UpdateUserPasswordPinTokenActivationStatusByUserIDParams{
		UserID: user.ID,
	}
	if user.Pin != "" {
		updateParam.Pin = pgtype.Text{String: user.Pin, Valid: true}
	}
	if user.Password != "" {
		updateParam.Password = pgtype.Text{String: user.Password, Valid: true}
	}
	if user.ActivationToken != "" {
		updateParam.TokenActivation = pgtype.Text{String: user.ActivationToken, Valid: true}
	}
	if !user.ActivationExpiredAt.IsZero() {
		updateParam.TokenActivationExpiredAt = pgtype.Timestamp{Time: user.ActivationExpiredAt, Valid: true}
	}
	if user.Status > 0 {
		updateParam.Status = pgtype.Int2{Int16: int16(user.Status), Valid: true}
	}
	_, err := q.UpdateUserPasswordPinTokenActivationStatusByUserID(ctx, updateParam)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (user domain.User, err error) {
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(r.GetUserByEmail))
	defer span.End()
	userDb, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, constant.ErrUserNotFound
		}
		return user, errors.Wrap(err, "get user by email")
	}
	err = utils.ObjectToObject(&userDb, &user)
	return user, err
}

func (r *Repository) InsertUserWithoutPassword(ctx context.Context, user domain.User, tx pgx.Tx) error {
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(r.InsertUserWithoutPassword))
	defer span.End()

	fullName := pgtype.Text{}
	q := r.q
	if tx != nil {
		q = sqlc.New(tx)
	}
	if user.FullName != "" {
		fullName = pgtype.Text{String: user.FullName, Valid: true}
	}
	_, err := q.InsertWithoutPassword(ctx, sqlc.InsertWithoutPasswordParams{
		Email:                    user.Email,
		FullName:                 fullName,
		Status:                   user.Status,
		RoleID:                   user.RoleID,
		TokenActivation:          pgtype.Text{String: user.ActivationToken, Valid: true},
		TokenActivationExpiredAt: pgtype.Timestamp{Time: user.ActivationExpiredAt, Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserById(ctx context.Context, id int64) (user domain.UserSimplified, err error) {
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(r.GetUserById))
	defer span.End()
	userDb, err := r.q.GetUserById(ctx, id)
	if err != nil {
		return user, err
	}
	err = utils.ObjectToObject(&userDb, &user)
	return user, err
}

func NewUserRepository(q *pgxpool.Pool) *Repository {
	return &Repository{q: sqlc.New(q)}
}
