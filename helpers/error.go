package helpers

import "errors"

var (
	ErrRecordNotFound  = errors.New("record not found")
	ErrInvalidRequest  = errors.New("invalid request data")
	ErrEmailIsVerified = errors.New("email is verified")
)
