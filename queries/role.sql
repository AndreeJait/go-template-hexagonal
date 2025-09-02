-- name: GetRoleById :one
SELECT id, name from roles where id = sqlc.arg(role_id);

-- name: GetRoles :many
SELECT id, name from roles OFFSET $1 LIMIT $2;