package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// @todo pull this into a service
func (h *Handler) Login(c echo.Context) (err error) {
	// diminish brute force attempts
	time.Sleep(500 * time.Millisecond)

	// Throws unauthorized error
	email := c.FormValue("email")
	p, err := h.App.RepositoryFactory.User().GetByEmail(email)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	} else if p.CheckPassword(c.FormValue("password")) == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		p.ID.String(),
		p.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(h.App.JwtSecureKey))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func (h *Handler) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return c.JSON(http.StatusOK, fmt.Sprintf("It Worked! %s!", claims.Email))
}
