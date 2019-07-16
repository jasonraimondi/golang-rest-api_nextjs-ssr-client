package lib

import (
	"database/sql"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/go-playground/validator.v9"

	"git.jasonraimondi.com/jason/jasontest/domain/repository"
	"git.jasonraimondi.com/jason/jasontest/domain/service"
)

type Application struct {
	//S3Config          *aws.Config
	//Validator         *validator.Validate
	RepositoryFactory *repository.Factory
	ServiceFactory    *service.Service
}

func NewApplication(dbx *sqlx.DB, s3Config *aws.Config) *Application {
	v := validator.New()
	_ = v.RegisterValidation("password-strength", ValidatePasswordStrength)
	r := repository.NewFactory(dbx)
	s := service.NewService(r, v, s3Config)
	return &Application{
		//Validator:         v,
		RepositoryFactory: r,
		ServiceFactory:    s,
	}
}

func NewTestApplication() (a *Application) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	dbx := sqlx.NewDb(db, "sqlite3")
	databaseInstance, err := sqlite3.WithInstance(dbx.DB, &sqlite3.Config{})
	if err != nil {
		panic(err)
	}
	sessionToken := ""
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials("", "", sessionToken),
		Endpoint:         aws.String("https://s3.wasabisys.com"),
		Region:           aws.String("us-east-1"),
		S3ForcePathStyle: aws.Bool(true),
	}
	a = NewApplication(dbx, s3Config)
	if err = repository.MigrateNow(&databaseInstance); err != nil {
		panic(err)
	}
	return a
}
