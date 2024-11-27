CREATE TABLE orders (
    order_id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(user_id),
    menu_id UUID REFERENCES menus(menu_id),
    status VARCHAR,
    payment_method PAYMENT_METHOD_ENUM,
    payment_proof_link VARCHAR(255) DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);