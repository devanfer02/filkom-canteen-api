CREATE TABLE admins (
    admin_id UUID PRIMARY KEY DEFAULT generate_ulid(),
    role_id UUID REFERENCES roles(role_id),
    fullname VARCHAR(150),
    wa_number VARCHAR(20),
    username VARCHAR(200) UNIQUE,
    password VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);