package domain

type Order struct {
	ID               string `json:"order_id" db:"order_id"`
	UserID           string `json:"user_id" db:"order_user_id"`
	MenuID           string `json:"menu_id" db:"order_menu_id"`
	Status           string `json:"status" db:"status"`
	PaymentMethod    string `json:"payment_method" db:"payment_method"`
	PaymentProofLink string `json:"payment_proof_link" db:"payment_proof_link"`
	CreatedAt        string `json:"created_at" db:"created_at"`
	UpdatedAt        string `json:"updated_at" db:"updated_at"`
}
