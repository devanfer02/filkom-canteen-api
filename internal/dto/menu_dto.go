package dto

import "mime/multipart"

type MenuParams struct {
	ID     string
	ShopID string
}

type MenuRequest struct {
	Name      string `json:"menu_name" db:"menu_name" binding:"required"`
	ShopID    string `json:"shop_id" db:"shop_id" binding:"required"`
	Price     int64  `json:"menu_price" db:"menu_price" binding:"required"`
	Status    string `json:"menu_status" db:"menu_status" binding:"required"`
	PhotoLink string `json:"menu_photo_link" db:"menu_photo_link"`
	Photo     *multipart.FileHeader
}
