package web

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func (h *Handler) Login(c echo.Context) (err error) {
	// diminish brute force attempts
	time.Sleep(500 * time.Millisecond)

	// Throws unauthorized error
	p, err := h.App.RepositoryFactory().Person().GetByEmail(c.FormValue("email"))

	// Set custom claims
	claims := &JwtCustomClaims{
		p.Email,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(h.JwtSecureKey))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func (h *Handler) Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func (h *Handler) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return c.JSON(http.StatusOK, fmt.Sprintf("It Worked! %s!", claims.Email))
}
