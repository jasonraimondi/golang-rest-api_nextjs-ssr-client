package web

import (
	"git.jasonraimondi.com/jason/jasontest/domain/lib"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/go-playground/validator.v9"
)

type Handler struct {
	App          *lib.Application
	JwtSecureKey string
	Validator    *validator.Validate
}

func NewHandler(a *lib.Application, j string) *Handler {
	v := validator.New()
	_ = v.RegisterValidation("password-strength", ValidatePasswordStrength)

	return &Handler{
		App:          a,
		JwtSecureKey: j,
		Validator:    v,
	}
}

func ValidatePasswordStrength(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) > 6
}

type JwtCustomClaims struct {
	Email      string `json:"email"`
	IsVerified bool   `json:"isVerified"`
	jwt.StandardClaims
}
