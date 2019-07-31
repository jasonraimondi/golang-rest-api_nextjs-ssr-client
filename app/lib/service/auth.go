package service

import (
	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type AuthService struct {
	userRepository *repository.UserRepository
	jwtSecureKey   string
}

func (s *AuthService) AttemptLogin(email string, password string) (string, *echo.HTTPError) {
	time.Sleep(500 * time.Millisecond)

	p, err := s.userRepository.GetByEmail(email)

	if err != nil {
		return "", echo.NewHTTPError(http.StatusNotFound, "user not found")
	} else if p.CheckPassword(password) == false {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
	}

	claims := &JwtCustomClaims{
		UserID: p.ID.String(),
		Email:  p.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(s.jwtSecureKey))
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "token error")
	}
	return t, nil
}

type JwtCustomClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
