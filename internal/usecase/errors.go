package usecase

import "errors"

var (
	ErrUserIsExists       = errors.New("user is exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
