package handlers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/dgrijalva/jwt-go"

	"git.jasonraimondi.com/jason/jasontest/lib"
)

type Handler struct {
	App      *lib.Application
	S3Config *aws.Config
}

func NewHandler(a *lib.Application) *Handler {
	return &Handler{
		App: a,
	}
}

type JwtCustomClaims struct {
	UserID string `json:"user_id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
