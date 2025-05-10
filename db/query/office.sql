-- name: CreateOffice :one
INSERT INTO offices (
  code,
  name,
  province_id,
  regency_id,
  district_id,
  email,
  phone,
  address,
  logo_url,
  created_by
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetOfficeById :one
SELECT  a.id, a.code, a.name, a.province_id, a.regency_id, a.district_id, a.email, a.phone, a.address, a.logo_url, a.created_by FROM offices a
WHERE a.id = $1 
AND a.deleted_at IS NULL 
LIMIT 1;

-- name: ListAllOffices :many
SELECT a.id, a.code, a.name, a.province_id, a.regency_id, a.district_id, a.email, a.phone, a.address, a.logo_url, a.created_by, b.name as province, c.name as regency, d.name as district FROM offices a
JOIN loc_provinces b ON a.province_id = b.id
JOIN loc_regencies c ON a.regency_id = c.id
LEFT JOIN loc_districts d ON a.district_id = d.id
WHERE a.deleted_at IS NULL 
ORDER BY a.id
LIMIT $1 OFFSET $2;

-- name: TotalListAllOffices :one
SELECT COUNT(*) as total_items FROM offices
WHERE deleted_at IS NULL;

-- name: UpdateOffice :one
UPDATE offices
SET code = $1,
  name = $2,
  province_id = $3,
  regency_id = $4,
  district_id = $5,
  email = $6,
  phone = $7,
  address = $8,
  logo_url = $9,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $10
AND deleted_at IS NULL
RETURNING *;

-- name: DeleteOffice :execresult
UPDATE offices
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
AND deleted_at IS NULL
RETURNING *;

-- name: ListOfficesWithFilters :many
SELECT a.id, a.code, a.name, a.province_id, a.regency_id, a.district_id, a.email, a.phone, a.address, a.logo_url, a.created_by, b.name as province, c.name as regency, d.name as district 
FROM offices a
JOIN loc_provinces b ON a.province_id = b.id
JOIN loc_regencies c ON a.regency_id = c.id
LEFT JOIN loc_districts d ON a.district_id = d.id
WHERE a.deleted_at IS NULL 
-- Optional province filter
AND (sqlc.narg(province_id)::int IS NULL OR a.province_id = sqlc.narg(province_id))
-- Optional regency filter
AND (sqlc.narg(regency_id)::int IS NULL OR a.regency_id = sqlc.narg(regency_id))
-- Optional district filter
AND (sqlc.narg(district_id)::int IS NULL OR a.district_id = sqlc.narg(district_id))
-- Optional name search
AND (sqlc.narg(name)::text IS NULL OR a.name ILIKE '%' || sqlc.narg(name) || '%')
ORDER BY a.id
LIMIT $1 OFFSET $2;

-- name: TotalListOfficesWithFilters :one
SELECT COUNT(*) as total_items 
FROM offices a
WHERE a.deleted_at IS NULL 
-- Optional province filter
AND (sqlc.narg(province_id)::int IS NULL OR a.province_id = sqlc.narg(province_id))
-- Optional regency filter
AND (sqlc.narg(regency_id)::int IS NULL OR a.regency_id = sqlc.narg(regency_id))
-- Optional district filter
AND (sqlc.narg(district_id)::int IS NULL OR a.district_id = sqlc.narg(district_id))
-- Optional name search
AND (sqlc.narg(name)::text IS NULL OR a.name ILIKE '%' || sqlc.narg(name) || '%');
