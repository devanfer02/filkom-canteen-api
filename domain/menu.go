package domain

import "time"

type Menu struct {
	ID        string    `json:"menu_id" db:"menu_id"`
	Name      string    `json:"menu_name" db:"menu_name"`
	ShopID    string    `json:"shop_id" db:"menu_shop_id"`
	Price     int64     `json:"price" db:"menu_price"`
	Status    string    `json:"status" db:"menu_status"`
	PhotoLink string    `json:"menu_photo_link" db:"menu_photo_link"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
