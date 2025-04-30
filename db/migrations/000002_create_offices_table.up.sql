CREATE TABLE offices (
    id BIGINT UNSIGNED PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    province_id INT NOT NULL,
    regency_id INT NOT NULL,
    district_id INT NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(100) NULL,
    address VARCHAR(255) NULL,
    logo_url VARCHAR(255) NULL,
    created_by BIGINT UNSIGNED NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL DEFAULT NULL,

    FOREIGN KEY (province_id) REFERENCES loc_provinces(id),
    FOREIGN KEY (regency_id) REFERENCES loc_regencies(id),
    FOREIGN KEY (district_id) REFERENCES loc_districts(id) 
);

-- NEED TO MAKE SEEDER TO OFFICES AND SCHOOLS