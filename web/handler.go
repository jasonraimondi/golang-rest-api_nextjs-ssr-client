package web

import (
	"git.jasonraimondi.com/jason/jasontest/domain/lib"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
)

type Handler struct {
	App *lib.Application
	JWT middleware.JWTConfig
}

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}