package service

import (
	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func (f *Service) CreateUser(email string, firstName string, lastName string, password string) (u *model.User, httpErr *echo.HTTPError) {
	if err := guardAgainstInvalidEmail(f.validate, email); err != nil {
		return nil, echo.NewHTTPError(http.StatusNotAcceptable, "invalid email", err)
	} else if err = guardAgainstDuplicateEmail(f.repository.User(), email); err != nil {
		return nil, echo.NewHTTPError(http.StatusConflict, "duplicate email", err)
	}

	u = model.NewUser(email)
	u.FirstName = model.ToNullString(firstName)
	u.LastName = model.ToNullString(lastName)

	if password != "" {
		if err := guardAgainstInvalidPassword(f.validate, password); err != nil {
			return nil, echo.NewHTTPError(http.StatusNotAcceptable, "invalid password", err)
		} else if err = u.SetPassword(password); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "server error set user password")
		}
	}
	return u, httpErr
}

func (f *Service) CreateSignUpConfirmation(u *model.User) (s *model.SignUpConfirmation, httpErr *echo.HTTPError) {
	s = model.NewSignUpConfirmation(*u)
	tx := f.repository.DBx.MustBegin()
	if err := repository.CreateUserTx(tx, *u); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err, "server error creating user", err)
	}
	repository.CreateSignUpConfirmationTx(tx, s)
	if err := tx.Commit(); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err, "server error transaction commit user", err)
	}
	return s, httpErr
}

func (f *Service) ValidateEmailSignUpConfirmation(token string, userId string) *echo.HTTPError {
	tx := f.repository.DBx.MustBegin()
	signUpConfirmation, err := repository.GetByTokenTx(tx, token)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "token not found")
	}
	user, err := repository.GetByIdTx(tx, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "error token not found")
	} else if err := tx.Commit(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error transaction failed", err)
	} else if signUpConfirmation.UserId.String() != userId {
		return echo.NewHTTPError(http.StatusNotAcceptable, "invalid user and token id")
	}
	user.SetVerified()
	tx = f.repository.DBx.MustBegin()
	repository.UpdateUserTx(tx, *user)
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
