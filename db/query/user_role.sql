-- name: CreateUserRole :one
INSERT INTO user_roles (
  id,
  name
) VALUES ($1, $2)
RETURNING *;

-- name: GetUserRoleById :one 
SELECT * FROM user_roles
WHERE id = $1 
AND deleted_at IS NULL
LIMIT 1;

-- name: DeleteUserRole :execresult
UPDATE user_roles 
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
AND deleted_at IS NULL
RETURNING *;

-- name: ListAllUserRoles :many
SELECT * FROM user_roles
WHERE deleted_at IS NULL 
ORDER BY id
LIMIT $1 OFFSET $2;