package handlers

import (
	"git.jasonraimondi.com/jason/jasontest/lib/service"
	"github.com/labstack/echo"
	"net/http"
)

type AuthHandler struct {
	factory *service.Factory
}

// @todo pull this into a service
func (h *AuthHandler) Login(c echo.Context) (err error) {
	token, httpErr := h.factory.AuthService().AttemptLogin(
		c.FormValue("email"),
		c.FormValue("password"),
	)
	if httpErr != nil {
		return httpErr;
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
