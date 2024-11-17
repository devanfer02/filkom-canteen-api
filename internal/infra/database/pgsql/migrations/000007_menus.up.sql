CREATE TABLE menus (
    menu_id UUID PRIMARY KEY DEFAULT generate_ulid(),
    menu_name VARCHAR(200) ,
    menu_price INTEGER,
    menu_status MENU_STATUS_ENUM,
    menu_photo_link VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);