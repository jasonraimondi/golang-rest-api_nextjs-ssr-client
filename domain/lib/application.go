package lib

import (
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
	"git.jasonraimondi.com/jason/jasontest/domain/service"
	"github.com/jmoiron/sqlx"
	"gopkg.in/go-playground/validator.v9"
)

type Application struct {
	Validator         *validator.Validate
	RepositoryFactory *repository.Factory
	ServiceFactory    *service.Service
}

func NewApplication(dbx *sqlx.DB) *Application {
	v := validator.New()
	_ = v.RegisterValidation("password-strength", ValidatePasswordStrength)
	r := repository.NewFactory(dbx)
	s := service.NewService(r, v)
	return &Application{
		Validator:         v,
		RepositoryFactory: r,
		ServiceFactory:    s,
	}
}
