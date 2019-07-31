package handlers

import (
	"git.jasonraimondi.com/jason/jasontest/lib"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/labstack/echo"
)

type Handler struct {
	App      *lib.Application
	S3Config *aws.Config
}

func NewHandler(a *lib.Application) *Handler {
	return &Handler{
		App: a,
	}
}

func (h *Handler) SignUp() *SignUpHandler {
	return &SignUpHandler{
		signUp: h.App.ServiceFactory.SignUpService(),
	}
}

func (h *Handler) Auth() *AuthHandler {
	return &AuthHandler{
		factory: h.App.ServiceFactory,
	}
}

func sendMessage(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, map[string]interface{}{
		"message": message,
	})
}
