package dto

import "mime/multipart"

type MenuParams struct {
	ID string
}

type MenuRequest struct {
	Name      string `json:"menu_name" db:"menu_name"`
	ShopID    string `json:"shop_id" db:"shop_id"`
	Price     int64  `json:"price" db:"price"`
	Status    string `json:"status" db:"status"`
	PhotoLink string `json:"menu_photo_link" db:"menu_photo_link"`
	Photo     *multipart.FileHeader 
}
