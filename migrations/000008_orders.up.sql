CREATE TABLE orders (
    order_id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(user_id),
    menu_id UUID REFERENCES menus(menu_id),
    status_id INTEGER REFERENCES order_status(status_id),
    payment_method PAYMENT_METHOD_ENUM,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);