package lib

import (
	"gopkg.in/go-playground/validator.v9"
)

func ValidatePasswordStrength(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) > 6
}
