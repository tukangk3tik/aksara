-- name: CreateOffice :execresult
INSERT INTO offices (
  id,
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
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetOffice :one
SELECT * FROM offices 
WHERE id = ? 
AND deleted_at IS NULL 
LIMIT 1;

-- name: ListAllOffices :many
SELECT * FROM offices
WHERE deleted_at IS NULL 
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListOfficesByProvince :many
SELECT * FROM offices
WHERE deleted_at IS NULL
AND province_id = ? 
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListOfficesByRegency :many
SELECT * FROM offices
WHERE deleted_at IS NULL
AND regency_id = ? 
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListOfficesByDistrict :many
SELECT * FROM offices
WHERE deleted_at IS NULL
AND district_id = ? 
ORDER BY id
LIMIT ? OFFSET ?;

-- name: UpdateOffice :execresult
UPDATE offices
SET code = ?,
  name = ?,
  province_id = ?,
  regency_id = ?,
  district_id = ?,
  email = ?,
  phone = ?,
  address = ?,
  logo_url = ?,
  updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteOffice :exec
UPDATE offices
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ?;
