package auth

import (
	"context"
	"github.com/AndreeJait/go-template-hexagonal/internal/domain"
	"github.com/jackc/pgx/v5"
)

type UserRepo interface {
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	InsertUserWithoutPassword(ctx context.Context, user domain.User, tx pgx.Tx) error
	UpdateUserPasswordPinTokenActivationStatusByUserIDParams(ctx context.Context, user domain.User, tx pgx.Tx) error
}

type RoleRepo interface {
	GetRoles(ctx context.Context, offset, limit int32) ([]domain.RoleSimplified, error)
	GetRoleByID(ctx context.Context, roleId int64) (domain.RoleSimplified, error)
}

type PermissionRepo interface {
	GetAllPermissionsByRoleID(ctx context.Context, roleId int64) ([]domain.PermissionSimplified, error)
}

type UseCase interface {
	GetUserByEmail(ctx context.Context, req GetUserByEmailRequest) (domain.User, error)
	CreateUser(ctx context.Context, req CreateUserRequest) error
	CreatePassword(ctx context.Context, req CreateUserPasswordRequest) (CreateUserPasswordResponse, error)
	CreatePin(ctx context.Context, req CreateUserPinRequest) error
	Login(ctx context.Context, req LoginRequest) (LoginResponse, error)
}
