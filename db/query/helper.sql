
-- name: ClearUsers :exec
TRUNCATE TABLE users RESTART IDENTITY CASCADE;

-- name: ClearUserRoles :exec
TRUNCATE TABLE user_roles RESTART IDENTITY CASCADE;

-- name: ClearOffices :exec
TRUNCATE TABLE offices RESTART IDENTITY CASCADE;

-- name: ClearSchools :exec
TRUNCATE TABLE schools RESTART IDENTITY CASCADE;

-- name: ClearLOVs :exec
TRUNCATE TABLE lovs RESTART IDENTITY CASCADE;