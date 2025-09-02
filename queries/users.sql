-- name: GetUserByEmail :one
SELECT
    id, email, full_name, status, password, pin, role_id, created_at, updated_at,
    token_activation, token_activation_expired_at
FROM users
WHERE email = $1;

-- name: GetUserById :one
SELECT
    id, email, full_name, status, password, pin, role_id, created_at
FROM users
WHERE id = $1;

-- name: InsertWithoutPassword :execrows
INSERT INTO users (email, full_name, status, role_id, token_activation,token_activation_expired_at)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateUserPasswordPinTokenActivationStatusByUserID :execrows
UPDATE users
SET
    password    = COALESCE(sqlc.narg(password),    password),
    pin         = COALESCE(sqlc.narg(pin),         pin),
    token_activation = COALESCE(sqlc.narg(token_activation), token_activation),
    token_activation_expired_at = COALESCE(sqlc.narg(token_activation_expired_at), token_activation_expired_at),
    status = COALESCE(sqlc.narg(status), status),
    updated_at       = now()
WHERE id = sqlc.arg(user_id);