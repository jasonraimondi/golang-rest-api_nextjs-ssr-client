package service

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"git.jasonraimondi.com/jason/jasontest/app/lib/config"
	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
)

type AuthService struct {
	userRepository *repository.UserRepository
	jwtSecureKey   config.JWTSecureKey
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
		UserID:      p.ID.String(),
		Email:       p.Email,
		IsVerified:  p.IsVerified,
		IsConfirmed: p.IsVerified, // @todo add is_confirmed column to user
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
	UserID      string `json:"userId"`
	Email       string `json:"email"`
	IsVerified  bool   `json:"isVerified"`
	IsConfirmed bool   `json:"isConfirmed"`
	jwt.StandardClaims
}
