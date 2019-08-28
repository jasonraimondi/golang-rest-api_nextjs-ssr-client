package service

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type SignUpService struct {
	validate                     *validator.Validate
	signUpConfirmationRepository *repository.SignUpConfirmationRepository
	userRepository               *repository.UserRepository
}

func (s *SignUpService) CreateUser(email string, firstName string, lastName string, password string) (u *models.User, httpErr *echo.HTTPError) {
	if err := guardAgainstInvalidEmail(s.validate, email); err != nil {
		return u, echo.NewHTTPError(http.StatusNotAcceptable, "invalid email", err)
	} else if err = guardAgainstDuplicateEmail(s.userRepository, email); err == nil {
		return u, echo.NewHTTPError(http.StatusConflict, "duplicate email", err)
	}

	u = models.NewUser(email)
	u.SetFirst(firstName)
	u.SetLast(lastName)

	if password != "" {
		if err := guardAgainstInvalidPassword(s.validate, password); err != nil {
			return nil, echo.NewHTTPError(http.StatusNotAcceptable, errors.New("invalid password"), err)
		} else if err = u.SetPassword(password); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, errors.New("server error set user password"), err)
		}
	}
	return u, httpErr
}

func (s *SignUpService) CreateSignUpConfirmation(u *models.User) (c *models.SignUpConfirmation, httpErr *echo.HTTPError) {
	c = models.NewSignUpConfirmation(u)
	if err := s.userRepository.Create(*u); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, errors.New("server error creating user"), err)
	}
	if err := s.signUpConfirmationRepository.Create(c); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, errors.New("server error creating sign up confirmation"), err)
	}
	return c, httpErr
}

// @todo return normal error, returning the echo http error is probably not a great idea, see sign_up_service_test
func (s *SignUpService) ValidateEmailSignUpConfirmation(token string, userId string) *echo.HTTPError {
	signUpConfirmation, err := s.signUpConfirmationRepository.GetByToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errors.New("token not found"))
	}
	user, err := s.userRepository.GetById(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errors.New("user not found"))
	}
	if signUpConfirmation.UserID.String() != userId {
		return echo.NewHTTPError(http.StatusNotAcceptable, errors.New("invalid user and token id"))
	}
	user.SetVerified()
	if err = s.userRepository.Update(*user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("error transaction failed"))
	} else if err = s.signUpConfirmationRepository.Delete(&signUpConfirmation); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("error transaction failed"))
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
