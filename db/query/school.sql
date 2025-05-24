-- name: CreateSchool :one
INSERT INTO schools (
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
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: GetSchoolById :one
SELECT a.id, a.code, a.name, a.office_id, a.province_id, a.regency_id, 
a.district_id, a.email, a.phone, a.address, a.logo_url, a.created_by, 
b.name as office, c.name as province, d.name as regency, e.name as district,
a.created_at, a.updated_at
FROM schools a
JOIN offices b ON a.office_id = b.id
JOIN loc_provinces c ON a.province_id = c.id
JOIN loc_regencies d ON a.regency_id = d.id
LEFT JOIN loc_districts e ON a.district_id = e.id
WHERE a.deleted_at IS NULL 
AND a.id = $1;

-- name: ListAllSchools :many
SELECT a.id, a.code, a.name, a.office_id, a.province_id, a.regency_id, a.district_id, a.email, a.phone, a.address, a.logo_url, a.created_by, b.name as office, c.name as province, d.name as regency, e.name as district 
FROM schools a
JOIN offices b ON a.office_id = b.id
JOIN loc_provinces c ON a.province_id = c.id
JOIN loc_regencies d ON a.regency_id = d.id
LEFT JOIN loc_districts e ON a.district_id = e.id
WHERE a.deleted_at IS NULL 
ORDER BY a.id
LIMIT $1 OFFSET $2;

-- name: TotalListAllSchools :one
SELECT COUNT(*) as total_items FROM schools
WHERE deleted_at IS NULL;

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
SET name = $1,
  office_id = $2,
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

-- name: DeleteSchool :execresult
UPDATE schools
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
AND deleted_at IS NULL
RETURNING *;
