-- name: CreateLov :one
INSERT INTO lovs (
  group_key,
  param_key,
  param_description,
  parent_id
) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetLovByParamKey :one
SELECT * FROM lovs
WHERE param_key = $1 
AND deleted_at IS NULL
LIMIT 1;

-- name: GetLovById :one
SELECT * FROM lovs
WHERE id = $1 
AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateLov :one
UPDATE lovs 
SET group_key = $1,
  param_key = $2,
  param_description = $3,
  parent_id = $4,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $5
AND deleted_at IS NULL
RETURNING *;

-- name: DeleteLov :execresult
UPDATE lovs 
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
AND deleted_at IS NULL
RETURNING *;

-- name: ListAllLovs :many
SELECT * FROM lovs
WHERE deleted_at IS NULL 
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: ListLovByGroupKey :many
SELECT * FROM lovs
WHERE group_key = $1
AND ((sqlc.narg(param_key)::text IS NULL OR param_key ILIKE '%' || sqlc.narg(param_key) || '%')
OR (sqlc.narg(param_description)::text IS NULL OR param_description ILIKE '%' || sqlc.narg(param_description) || '%'))
AND deleted_at IS NULL
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: TotalListAllLovs :one
SELECT COUNT(*) as total_items FROM lovs
WHERE deleted_at IS NULL;

-- name: TotalListLovByGroupKey :one
SELECT COUNT(*) as total_items FROM lovs
WHERE group_key = $1 
AND ((sqlc.narg(param_key)::text IS NULL OR param_key ILIKE '%' || sqlc.narg(param_key) || '%')
OR (sqlc.narg(param_description)::text IS NULL OR param_description ILIKE '%' || sqlc.narg(param_description) || '%'))
AND deleted_at IS NULL;