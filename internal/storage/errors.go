package storage

import "errors"

var (
	ErrEmailIsExists = errors.New("email is exists")
	ErrEmailNotFound = errors.New("email not found")
)
