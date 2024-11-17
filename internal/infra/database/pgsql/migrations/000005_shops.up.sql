CREATE TABLE shops (
    shop_id UUID PRIMARY KEY DEFAULT generate_ulid(),
    shop_name VARCHAR(200),
    shop_description TEXT,
    shop_foto_link VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);