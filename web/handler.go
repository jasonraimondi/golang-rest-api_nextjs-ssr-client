package web

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/dgrijalva/jwt-go"

	"git.jasonraimondi.com/jason/jasontest/domain/lib"
)

// struct is kind of like "class" or object in javascript, add methods to the struct
//in different files
type Handler struct {
	App          *lib.Application
	S3Config     *aws.Config
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
