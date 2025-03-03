package controller

import "errors"

var (
	ErrBadRequest         = errors.New("bad request")
	ErrFailedValidate     = errors.New("failed validation")
	ErrInternalError      = errors.New("internal server error")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
