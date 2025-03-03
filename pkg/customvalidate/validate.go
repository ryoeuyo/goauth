package customvalidate

import (
	"github.com/go-playground/validator/v10"
)

func IsValidPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	return len(password) >= 8 && len(password) <= 32
}
