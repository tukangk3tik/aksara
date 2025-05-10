-- name: CreateUser :one
INSERT INTO users (
  name,
  fullname,
  email,
  password,
  user_role_id,
  office_id,
  school_id
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateUser :one
UPDATE users 
SET name = $1,
  fullname = $2,
  email = $3,
  password = $4,
  user_role_id = $5,
  office_id = $6,
  school_id = $7,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $8
AND deleted_at IS NULL
RETURNING *;

-- name: ListAllUsers :many
SELECT * FROM users
WHERE deleted_at IS NULL 
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: TotalListAllUsers :one
SELECT COUNT(*) as total_items FROM users
WHERE deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 
AND deleted_at IS NULL
LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 
AND deleted_at IS NULL
LIMIT 1;

-- name: DeleteUser :execresult
UPDATE users 
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
AND deleted_at IS NULL
RETURNING *;