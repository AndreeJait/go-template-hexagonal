-- name: GetAllPermissionsByRoleID :many
SELECT
    p.id, p.name
FROM role_permissions rp
JOIN permissions p ON p.id = rp.permission_id
WHERE rp.role_id = sqlc.arg(role_id);