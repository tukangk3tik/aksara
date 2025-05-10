-- name: CreateSchool :one
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
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: GetSchoolById :one
SELECT * FROM schools 
WHERE id = $1 
AND deleted_at IS NULL 
LIMIT 1;

-- name: ListAllSchools :many
SELECT * FROM schools
WHERE deleted_at IS NULL 
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: ListSchoolsByProvince :many
SELECT * FROM schools
WHERE deleted_at IS NULL
AND province_id = $1 
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: ListSchoolsByRegency :many
SELECT * FROM schools
WHERE deleted_at IS NULL
AND regency_id = $1 
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: ListSchoolsByDistrict :many
SELECT * FROM schools
WHERE deleted_at IS NULL
AND district_id = $1 
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: ListSchoolsByOffice :many
SELECT * FROM schools
WHERE deleted_at IS NULL
AND office_id = $1 
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: UpdateSchool :one
UPDATE schools
SET code = $1,
  name = $2,
  office_id = $3,
  province_id = $4,
  regency_id = $5,
  district_id = $6,
  email = $7,
  phone = $8,
  address = $9,
  logo_url = $10,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $11  
AND deleted_at IS NULL
RETURNING *;

-- name: DeleteSchool :execresult
UPDATE schools
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
AND deleted_at IS NULL
RETURNING *;
