
-- name: ClearUsers :exec
TRUNCATE TABLE users RESTART IDENTITY CASCADE;

-- name: ClearUserRoles :exec
TRUNCATE TABLE user_roles RESTART IDENTITY CASCADE;