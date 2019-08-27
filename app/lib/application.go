package lib

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"

	"git.jasonraimondi.com/jason/jasontest/app/lib/awsupload"
	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
)

type Application struct {
	MigrationDir      string
	RepositoryFactory *repository.Factory
	ServiceFactory    *service.Factory
}

func NewApplication(dbx *gorm.DB, s3Config *awsupload.S3Config, jwtSecureKey string, dir string) *Application {
	validate := validator.New()
	_ = validate.RegisterValidation("password-strength", ValidatePasswordStrength)
	repositoryFactory := repository.NewFactory(dbx)
	serviceFactory := service.NewService(repositoryFactory, validate, s3Config, jwtSecureKey)
	return &Application{
		RepositoryFactory: repositoryFactory,
		ServiceFactory:    serviceFactory,
		MigrationDir:      dir,
	}
}
