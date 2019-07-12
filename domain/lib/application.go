package lib

import (
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
	"github.com/jmoiron/sqlx"
)

type Application struct {
	DBx *sqlx.DB
}

func NewApplication(dbx *sqlx.DB) *Application {
	return &Application{DBx: dbx}
}

func (a *Application) RepositoryFactory() *repository.Factory {
	return repository.NewFactory(a.DBx)
}
