package domain

import "errors"

var (
	ErrNotFound       = errors.New("item not found")
	ErrBadRequest     = errors.New("bad data request")
	ErrDuplicateEntry = errors.New("duplicate item entry")
	ErrInvalidToken   = errors.New("invalid token")
)

func GetStatus(err error) (int, string) {
	if err == nil {
		return 200, "success"
	}

	switch err {
	case ErrNotFound:
		return 404, "fail"
	case ErrBadRequest, ErrInvalidToken:
		return 400, "fail"
	case ErrDuplicateEntry:
		return 409, "fail"
	default:
		return 500, "error"
	}
}
