CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    fullname VARCHAR(100) NOT NULL, 
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    user_role_id INTEGER NOT NULL, 
    office_id BIGINT DEFAULT NULL,
    school_id BIGINT DEFAULT NULL,
    is_super_admin BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    FOREIGN KEY (user_role_id) REFERENCES user_roles(id),
    FOREIGN KEY (office_id) REFERENCES offices(id),
    FOREIGN KEY (school_id) REFERENCES schools(id)
);