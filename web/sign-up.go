package web

import (
	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func (h *Handler) ConfirmEmail(c echo.Context) (err error) {
	return c.JSON(http.StatusAccepted, http.StatusText(http.StatusAccepted))
}

func (h *Handler) SignUp(c echo.Context) (err error) {
	r := h.App.RepositoryFactory()
	email := c.FormValue("email")
	r.User()
	if err = guardAgainstInvalidEmail(h.Validator, email); err != nil {
		return err
	} else if err = guardAgainstDuplicateEmail(r.User(), email); err != nil {
		return err
	}

	u := model.NewUser(email)
	u.FirstName = model.ToNullString(c.FormValue("first_name"))
	u.LastName = model.ToNullString(c.FormValue("last_name"))

	password := c.FormValue("password")
	if password != "" {
		if err = guardAgainstInvalidPassword(h.Validator, password); err != nil {
			return err
		} else if err = u.SetPassword(password); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "server error set user password")
		}
	}
	t := model.NewSignUpConfirmation(u)
	tx := r.DBx.MustBegin()
	if err = r.User().CreateTx(tx, u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err, "server error creating user", err)
	}
	tx.MustExec(
		"INSERT INTO sign_up_confirmation (token, user_id, created_at) VALUES ($1, $2, $3)",
		t.Token,
		t.User.ID,
		t.CreatedAt,
	)
	if err = tx.Commit(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err, "server error transaction commit user", err)
	}
	return c.JSON(http.StatusCreated, http.StatusText(http.StatusCreated))
}

func guardAgainstInvalidPassword(v *validator.Validate, email string) (err error) {
	if err = v.Var(email, "required,password-strength"); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, "invalid password", err)
	}
	return nil
}

func guardAgainstInvalidEmail(v *validator.Validate, email string) (err error) {
	if err = v.Var(email, "required,email"); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, "invalid email", err)
	}
	return nil
}

func guardAgainstDuplicateEmail(r *repository.SqlxUserRepository, email string) (err error) {
	if _, err = r.GetByEmail(email); err == nil {
		return echo.NewHTTPError(http.StatusConflict, "duplicate email", err)
	}
	return nil
}