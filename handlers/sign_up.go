package handlers

import (
	"git.jasonraimondi.com/jason/jasontest/lib/service"
	"net/http"

	"github.com/labstack/echo"
)

type SignUpHandler struct {
	signUp *service.SignUpService
}

func (h *SignUpHandler) SignUp(c echo.Context) error {
	email := c.FormValue("email")
	firstName := c.FormValue("first")
	lastName := c.FormValue("last")
	password := c.FormValue("password")

	if user, httpError := h.signUp.CreateUser(email, firstName, lastName, password); httpError != nil {
		return httpError
	} else if _, httpError = h.signUp.CreateSignUpConfirmation(user); httpError != nil {
		return httpError
	}

	return sendMessage(c, http.StatusCreated, http.StatusText(http.StatusCreated))
}

func (h *SignUpHandler) SignUpConfirmation(c echo.Context) error {
	token := c.QueryParam("t")
	userId := c.QueryParam("u")

	if httpErr := h.signUp.ValidateEmailSignUpConfirmation(token, userId); httpErr != nil {
		return httpErr
	}

	return sendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}
