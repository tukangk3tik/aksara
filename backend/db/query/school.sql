-- name: CreateSchool :execresult
INSERT INTO schools (
  id,
  code,
  name,
  office_id,
  province_id,
  regency_id,
  district_id,
  email,
  phone,
  address,
  logo_url,
  created_by
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetSchool :one
SELECT * FROM schools 
WHERE id = ? 
AND deleted_at IS NULL 
LIMIT 1;

-- name: ListAllSchools :many
SELECT * FROM schools
WHERE deleted_at IS NULL 
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListSchoolsByProvince :many
SELECT * FROM schools
WHERE deleted_at IS NULL
AND province_id = ? 
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListSchoolsByRegency :many
SELECT * FROM schools
WHERE deleted_at IS NULL
AND regency_id = ? 
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListSchoolsByDistrict :many
SELECT * FROM schools
WHERE deleted_at IS NULL
AND district_id = ? 
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListSchoolsByOffice :many
SELECT * FROM schools
WHERE deleted_at IS NULL
AND office_id = ? 
ORDER BY id
LIMIT ? OFFSET ?;

-- name: UpdateSchool :execresult
UPDATE schools
SET code = ?,
  name = ?,
  office_id = ?,
  province_id = ?,
  regency_id = ?,
  district_id = ?,
  email = ?,
  phone = ?,
  address = ?,
  logo_url = ?,
  updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteSchool :exec
UPDATE schools
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ?;
