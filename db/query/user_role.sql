-- name: CreateUserRole :one
INSERT INTO user_roles (
  name
) VALUES ($1)
RETURNING *;

-- name: GetUserRoleById :one 
SELECT * FROM user_roles
WHERE id = $1 LIMIT 1;

-- name: DeleteUserRole :execresult
UPDATE user_roles 
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
AND deleted_at IS NULL
RETURNING *;