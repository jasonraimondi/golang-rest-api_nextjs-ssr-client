package web

import (
	"git.jasonraimondi.com/jason/jasontest/domain/lib"
	"github.com/dgrijalva/jwt-go"
)

type Handler struct {
	App          *lib.Application
	JwtSecureKey string
}

func NewHandler(a *lib.Application, j string) *Handler {
	return &Handler{
		App:          a,
		JwtSecureKey: j,
	}
}

type JwtCustomClaims struct {
	Email      string `json:"email"`
	IsVerified bool   `json:"isVerified"`
	jwt.StandardClaims
}
