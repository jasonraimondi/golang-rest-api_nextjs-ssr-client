package lib

import (
	"database/sql"
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
	"git.jasonraimondi.com/jason/jasontest/domain/service"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
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
	a = NewApplication(dbx)
	if err = repository.MigrateNow(&databaseInstance); err != nil {
		panic(err)
	}
	return a
}
