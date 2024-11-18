package domain

import "time"

type Shop struct {
	ID          string    `json:"shop_id" db:"shop_id"`
	Name        string    `json:"shop_name" db:"shop_name"`
	Description string    `json:"shop_description" db:"shop_description"`
	PhotoLink   string    `json:"shop_photo_link" db:"shop_photo_link"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
