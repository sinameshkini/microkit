package helpers

import (
	"errors"
)

var (
	ErrAlreadyExist          = errors.New("already exist")
	ErrNotFound              = errors.New("not found")
	ErrRecordNotFound        = errors.New("record not found")
	ErrInvalidRequest        = errors.New("invalid request data")
	ErrEmailIsVerified       = errors.New("email is verified")
	ErrInvalidRestClientInfo = errors.New("invalid rest client info")
	ErrCacheInfo             = errors.New("invalid cache info")
)
