-- name: LocationProvince :many
SELECT * FROM loc_provinces;

-- name: LocationRegencyByProvince :many
SELECT * FROM loc_regencies 
WHERE province_id = ?;

-- name: LocationDistrictByRegency :many
SELECT * FROM loc_districts
WHERE regency_id = ?;