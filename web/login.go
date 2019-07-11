package web

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func (h *Handler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// diminish brute force attempts
	time.Sleep(500 * time.Millisecond)

	// Throws unauthorized error
	if username != "jon" && password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		username,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("*JwtSecureKey"))
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
	return c.JSON(http.StatusOK, "It Worked! "+claims.Name+"!")
}
