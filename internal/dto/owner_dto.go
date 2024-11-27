package dto

type OwnerParams struct {
	ID string
}

type OwnerRequest struct {
	Fullname string `json:"fullname" db:"fullname"`
	WANumber string `json:"wa_number" db:"wa_number"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
