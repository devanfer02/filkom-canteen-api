package domain

import "errors"

var (
	ErrNotFound = errors.New("item not found")
	ErrBadRequest = errors.New("bad data request")
)

func GetStatus(err error) (int, string) {
	if err == nil {
		return 200, "success"
	}

	switch err {
	case ErrNotFound:
		return 404, "fail"
	case ErrBadRequest:
		return 400, "fail"
	default:
		return 500, "error"
	}
}