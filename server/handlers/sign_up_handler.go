package handlers

import (
	"net/http"

	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
	"git.jasonraimondi.com/jason/jasontest/server/responses"

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

	if user, err := h.signUp.CreateUser(email, firstName, lastName, password); err != nil {
		return err
	} else if _, err = h.signUp.CreateSignUpConfirmation(user); err != nil {
		return err
	}

	return responses.SendMessage(c, http.StatusCreated, http.StatusText(http.StatusCreated))
}

func (h *SignUpHandler) SignUpConfirmation(c echo.Context) error {
	token := c.QueryParam("t")
	userId := c.QueryParam("u")

	if err := h.signUp.ValidateEmailSignUpConfirmation(token, userId); err != nil {
		return err
	}

	return responses.SendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}
