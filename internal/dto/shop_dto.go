package dto

import "mime/multipart"

type ShopParams struct {
	ID      string
	OwnerID string
}

type ShopRequest struct {
	Name        string `json:"shop_name" binding:"required"`
	Description string `json:"shop_description" binding:"required"`
	Photo       *multipart.FileHeader
	PhotoLink   string `json:"photo_link"`
}
