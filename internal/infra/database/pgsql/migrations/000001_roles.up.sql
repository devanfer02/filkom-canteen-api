CREATE TABLE roles (
    role_id UUID PRIMARY KEY DEFAULT generate_ulid(),
    role_name VARCHAR(50)
);