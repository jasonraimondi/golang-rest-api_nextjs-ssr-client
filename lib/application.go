package lib

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"gopkg.in/go-playground/validator.v9"

	"git.jasonraimondi.com/jason/jasontest/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/lib/s3"
	"git.jasonraimondi.com/jason/jasontest/lib/service"
)

type Application struct {
	JwtSecureKey      string
	MigrationDir      string
	RepositoryFactory *repository.Factory
	ServiceFactory    *service.Factory
}

func NewApplication(dbx *sqlx.DB, s3Config *s3.S3Config, jwtSecureKey string, dir string) *Application {
	v := validator.New()
	_ = v.RegisterValidation("password-strength", ValidatePasswordStrength)
	r := repository.NewFactory(dbx)
	s := service.NewService(r, v, s3Config)
	return &Application{
		RepositoryFactory: r,
		ServiceFactory:    s,
		JwtSecureKey:      jwtSecureKey,
		MigrationDir:      dir,
	}
}
