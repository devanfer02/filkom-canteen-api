package domain

import "time"

type Owner struct {
	ID        string    `json:"owner_id" db:"admin_id"`
	Fullname  string    `json:"fullname" db:"fullname"`
	WANumber  string    `json:"wa_number" db:"wa_number"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
