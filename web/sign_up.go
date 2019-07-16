package web

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) ConfirmEmail(c echo.Context) error {
	r := h.App.ServiceFactory
	token := c.QueryParam("t")
	userId := c.QueryParam("u")
	if httpErr := r.ValidateEmailSignUpConfirmation(token, userId); httpErr != nil {
		return httpErr
	}
	return sendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}

func (h *Handler) SignUp(c echo.Context) error {
	s := h.App.ServiceFactory

	email := c.FormValue("email")
	firstName := c.FormValue("first")
	lastName := c.FormValue("last")
	password := c.FormValue("password")

	if user, httpError := s.CreateUser(email, firstName, lastName, password); httpError != nil {
		return httpError
	} else if _, httpError = s.CreateSignUpConfirmation(user); httpError != nil {
		return httpError
	}
	return sendMessage(c, http.StatusCreated, http.StatusText(http.StatusCreated))
}
