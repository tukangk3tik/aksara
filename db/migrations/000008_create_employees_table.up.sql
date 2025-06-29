CREATE TABLE employees (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT,
  fullname VARCHAR(255),
  resident_identity_number VARCHAR(255),
  employee_identity_number VARCHAR(255),
  gender smallint, -- 1 / 0
  birthdate date,
  address VARCHAR(255),
  religion int,
  birthday date,
  birth_place VARCHAR(255),
  photo VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL DEFAULT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id)
);