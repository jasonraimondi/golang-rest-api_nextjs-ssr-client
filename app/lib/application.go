package lib

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"

	"git.jasonraimondi.com/jason/jasontest/app/lib/config"
	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
)

type Application struct {
	RepositoryFactory *repository.Factory
	ServiceFactory    *service.Factory
}

func NewApplication(db *gorm.DB, s3Config *config.S3Config, jwtSecureKey config.JWTSecureKey, debug bool) *Application {
	validate := validator.New()
	_ = validate.RegisterValidation("password-strength", ValidatePasswordStrength)
	repositoryFactory := repository.NewFactory(db, debug)
	serviceFactory := service.NewService(repositoryFactory, validate, s3Config, jwtSecureKey)
	return &Application{
		RepositoryFactory: repositoryFactory,
		ServiceFactory:    serviceFactory,
	}
}
