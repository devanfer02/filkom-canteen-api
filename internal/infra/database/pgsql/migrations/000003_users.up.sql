CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT generate_ulid(),
    fullname VARCHAR(150),
    email VARCHAR(200),
    password VARCHAR(255),
    wa_number VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);