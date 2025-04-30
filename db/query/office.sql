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
SELECT  a.id, a.code, a.name, a.province_id, a.regency_id, a.district_id, a.email, a.phone, a.address, a.logo_url, a.created_by FROM offices a
WHERE a.id = ? 
AND a.deleted_at IS NULL 
LIMIT 1;

-- name: ListAllOffices :many
SELECT a.id, a.code, a.name, a.province_id, a.regency_id, a.district_id, a.email, a.phone, a.address, a.logo_url, a.created_by, b.name as province, c.name as regency, d.name as district FROM offices a
JOIN loc_provinces b ON a.province_id = b.id
JOIN loc_regencies c ON a.regency_id = c.id
LEFT JOIN loc_districts d ON a.district_id = d.id
WHERE a.deleted_at IS NULL 
ORDER BY a.id
LIMIT ? OFFSET ?;

-- name: TotalListAllOffices :one
SELECT COUNT(*) as total_items FROM offices
WHERE deleted_at IS NULL;

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
WHERE id = ?
AND deleted_at IS NULL;

-- name: DeleteOffice :execresult
UPDATE offices
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ?
AND deleted_at IS NULL;
