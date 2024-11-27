package dto

import "mime/multipart"

type OrderParams struct {
	ID     string
	UserID string
	MenuID string
}

type OrderRequest struct {
	ID               string                `json:"order_id"`
	UserID           string                `json:"user_id"`
	MenuID           string                `json:"menu_id"`
	Status           string                `json:"status"`
	PaymentMethod    string                `json:"payment_method"`
	PaymentProofLink string                `json:"payment_proof_link"`
	PaymentProofFile *multipart.FileHeader `form:"payment_proof"`
	CreatedAt        string                `json:"created_at"`
	UpdatedAt        string                `json:"updated_at"`
}
