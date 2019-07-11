package web

import (
	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) SignUp(c echo.Context) (err error) {
	p := model.NewPerson(c.FormValue("email"))
	p.FirstName = model.ToNullString(c.FormValue("first_name"))
	p.LastName = model.ToNullString(c.FormValue("last_name"))
	if err = p.SetPassword(c.FormValue("password")); err != nil {
		return err
	}
	if err = h.App.RepositoryFactory().Person().Create(p); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, http.StatusText(http.StatusCreated))
}
