package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
)

type AuthHandler struct {
	factory *service.Factory
}

func (h *AuthHandler) Login(c echo.Context) (err error) {
	token, httpErr := h.factory.AuthService().AttemptLogin(
		c.FormValue("email"),
		c.FormValue("password"),
	)
	if httpErr != nil {
		return httpErr
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
