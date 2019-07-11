package web

import (
	"git.jasonraimondi.com/jason/jasontest/domain/action"
	"git.jasonraimondi.com/jason/jasontest/domain/action_handlers"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) SignUp(c echo.Context) error {
	first := c.FormValue("first")
	last := c.FormValue("last")
	password := c.FormValue("password")

	r := h.App.RepositoryFactory()
	createHandler := action_handlers.NewCreatePersonHandler(r.Person())

	createPerson := action.NewCreatePerson(
		&first,
		&last,
		c.FormValue("email"),
		&password,
	)
	if err := createHandler.Handle(createPerson); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, createPerson)
}
