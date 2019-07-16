package service

import (
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
)

func (s *Service) CreateUser(email string, firstName string, lastName string, password string) (u *model.User, httpErr *echo.HTTPError) {
	if err := guardAgainstInvalidEmail(s.validate, email); err != nil {
		return nil, echo.NewHTTPError(http.StatusNotAcceptable, "invalid email", err)
	} else if err = guardAgainstDuplicateEmail(s.repository.User(), email); err == nil {
		return nil, echo.NewHTTPError(http.StatusConflict, "duplicate email", err)
	}

	u = model.NewUser(email)
	u.First = model.ToNullString(firstName)
	u.Last = model.ToNullString(lastName)

	if password != "" {
		if err := guardAgainstInvalidPassword(s.validate, password); err != nil {
			return nil, echo.NewHTTPError(http.StatusNotAcceptable, "invalid password", err)
		} else if err = u.SetPassword(password); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "server error set user password")
		}
	}
	return u, httpErr
}

func (s *Service) CreateSignUpConfirmation(u *model.User) (c *model.SignUpConfirmation, httpErr *echo.HTTPError) {
	c = model.NewSignUpConfirmation(*u)
	tx := s.repository.DBx.MustBegin()
	if err := repository.CreateUserTx(tx, u); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err, "server error creating user", err)
	}
	repository.CreateSignUpConfirmationTx(tx, c)
	if err := tx.Commit(); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err, "server error transaction commit user", err)
	}
	return c, httpErr
}

func (s *Service) ValidateEmailSignUpConfirmation(token string, userId string) *echo.HTTPError {
	tx := s.repository.DBx.MustBegin()
	signUpConfirmation, err := repository.GetByTokenTx(tx, token)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "token not found")
	}
	user, err := repository.GetByIdTx(tx, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	} else if err := tx.Commit(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error transaction failed", err)
	} else if signUpConfirmation.UserId.String() != userId {
		return echo.NewHTTPError(http.StatusNotAcceptable, "invalid user and token id")
	}
	user.SetVerified()
	tx = s.repository.DBx.MustBegin()
	repository.UpdateUserTx(tx, user)
	repository.DeleteSignUpConfirmationTx(tx, signUpConfirmation)
	if err := tx.Commit(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error transaction failed")
	}
	return nil
}

func guardAgainstInvalidPassword(v *validator.Validate, email string) (err error) {
	return v.Var(email, "required,password-strength")
}

func guardAgainstInvalidEmail(v *validator.Validate, email string) (err error) {
	return v.Var(email, "required,email")
}

func guardAgainstDuplicateEmail(r *repository.UserRepository, email string) (err error) {
	_, err = r.GetByEmail(email)
	return err
}
