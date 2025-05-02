-- name: LocationProvince :many
SELECT * FROM loc_provinces 
WHERE name LIKE $1
LIMIT $2 OFFSET $3;

-- name: LocationRegencyByProvince :many
SELECT * FROM loc_regencies 
WHERE province_id = $1 
AND name LIKE $2
LIMIT $3 OFFSET $4;

-- name: LocationDistrictByRegency :many
SELECT * FROM loc_districts
WHERE regency_id = $1 
AND name LIKE $2
LIMIT $3 OFFSET $4;