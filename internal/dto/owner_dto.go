package dto

type OwnerParams struct {
	ID string
}

type OwnerRequest struct {
	Fullname string `json:"fullname" db:"fullname" binding:"required"`
	WANumber string `json:"wa_number" db:"wa_number" binding:"required"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
