package web

import (
	"git.jasonraimondi.com/jason/jasontest/domain/lib"
	"github.com/dgrijalva/jwt-go"
)

type Handler struct {
	App          *lib.Application
	JwtSecureKey string
}

func NewHandler(a *lib.Application, jwtSecureKey string) *Handler {
	return &Handler{
		App:          a,
		JwtSecureKey: jwtSecureKey,
	}
}

func NewTestHandler() *Handler {
	a := lib.NewTestApplication()
	return &Handler{
		App:          a,
		JwtSecureKey: "testing",
	}
}

type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
