package dto

import "mime/multipart"

type ShopParams struct {
	ID      string
	OwnerID string
}

type ShopRequest struct {
	Name        string `json:"shop_name"`
	Description string `json:"shop_description"`
	Photo       *multipart.FileHeader
	PhotoLink   string
}
