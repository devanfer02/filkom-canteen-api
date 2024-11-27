package domain

import "time"

type User struct {
	ID        string    `json:"user_id" db:"user_id"`
	Fullname  string    `json:"fullname" db:"fullname"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	WANumber  string    `json:"wa_number" db:"wa_number"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
