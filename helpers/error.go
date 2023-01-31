package helpers

import (
	"errors"
	"net/http"
)

var (
	ErrAlreadyExist    = errors.New("already exist")
	ErrNotFound        = errors.New("not found")
	ErrRecordNotFound  = errors.New("record not found")
	ErrInvalidRequest  = errors.New("invalid request data")
	ErrEmailIsVerified = errors.New("email is verified")
)

func ErrorToHttpStatusCode(err error) int {
	switch err {
	case ErrNotFound, ErrRecordNotFound:
		return http.StatusNotFound

	default:
		return http.StatusInternalServerError
	}
}
