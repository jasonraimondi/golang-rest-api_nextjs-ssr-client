package lib

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	v := validator.New()
	_ = v.RegisterValidation("password-strength", ValidatePasswordStrength)
	r := repository.NewFactory(dbx)
	s := service.NewService(r, v, s3Config, jwtSecureKey)
	return &Application{
		RepositoryFactory: r,
		ServiceFactory:    s,
		MigrationDir:      dir,
	}
}
