-- name: CreateUser :execresult
INSERT INTO users (
  id,
  name,
  fullname,
  email,
  password,
  user_role_id,
  office_id,
  school_id
) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetUserByEmail :one
SELECT id, name, email, password FROM users
WHERE email = ? LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: DeleteUser :exec
UPDATE users 
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ?;