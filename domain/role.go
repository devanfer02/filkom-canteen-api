package domain

type Role struct {
	ID   string `db:"role_id"`
	Name string `db:"role_name"`
}
