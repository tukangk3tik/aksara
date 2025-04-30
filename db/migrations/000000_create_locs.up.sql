CREATE TABLE loc_provinces (
    id INT PRIMARY KEY,
    name varchar(100) NOT NULL
);

CREATE TABLE loc_regencies (
    id INT PRIMARY KEY,
    name varchar(100) NOT NULL,
    province_id INT NOT NULL,
    FOREIGN KEY (province_id) REFERENCES loc_provinces(id) 
);

CREATE TABLE loc_districts (
    id INT PRIMARY KEY,
    name varchar(100) NOT NULL,
    province_id INT NOT NULL,
    regency_id INT NOT NULL,
    FOREIGN KEY (province_id) REFERENCES loc_provinces(id),
    FOREIGN KEY (regency_id) REFERENCES loc_regencies(id) 
);