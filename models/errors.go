package models

import "errors"

var (
	ErrNotfound         = errors.New("not found")
	ErrAlreadyExist     = errors.New("already exist")
	ErrInvalidRequest   = errors.New("invalid request data")
	ErrInvalidID        = errors.New("id is not valid")
	ErrInternal         = errors.New("internal error, try again later")
	ErrPermissionDenied = errors.New("permission denied")

	ErrInvalidRestClientInfo = errors.New("invalid rest client info")
	ErrCacheInfo             = errors.New("invalid cache info")
	ErrConnectionLost        = errors.New("database connection lost")
	ErrInvalidConnection     = errors.New("invalid database connection")
	ErrInvalidDriver         = errors.New("invalid driver selected")
)
