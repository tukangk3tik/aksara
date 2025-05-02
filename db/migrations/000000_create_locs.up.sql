CREATE TABLE loc_provinces (
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL
);

CREATE TABLE loc_regencies (
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL,
    province_id INTEGER NOT NULL,
    FOREIGN KEY (province_id) REFERENCES loc_provinces(id) 
);

CREATE TABLE loc_districts (
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL,
    province_id INTEGER NOT NULL,
    regency_id INTEGER NOT NULL,
    FOREIGN KEY (province_id) REFERENCES loc_provinces(id),
    FOREIGN KEY (regency_id) REFERENCES loc_regencies(id) 
);