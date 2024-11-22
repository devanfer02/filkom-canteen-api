CREATE TABLE shop_owners (
    shop_id UUID REFERENCES shops(shop_id),
    admin_id UUID REFERENCES admins(admin_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);