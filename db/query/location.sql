-- name: LocationProvince :many
SELECT * FROM loc_provinces 
WHERE name LIKE ?
LIMIT ? OFFSET ?;

-- name: LocationRegencyByProvince :many
SELECT * FROM loc_regencies 
WHERE province_id = ? 
AND name LIKE ?
LIMIT ? OFFSET ?;

-- name: LocationDistrictByRegency :many
SELECT * FROM loc_districts
WHERE regency_id = ? 
AND name LIKE ?
LIMIT ? OFFSET ?;