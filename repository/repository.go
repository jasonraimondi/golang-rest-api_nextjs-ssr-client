package repository

import (
	"git.jasonraimondi.com/jason/jasontest/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type RepositoryFactory struct {
	DB *sqlx.DB
}

func (r RepositoryFactory) User() *PersonRepository {
	return &PersonRepository{db: r.DB}
}

func (r RepositoryFactory) Migrate(schema string) error {
	_ = r.DB.MustExec(schema)
	return nil
}

func (r RepositoryFactory) Seed() (err error) {
	u := models.NewSimplePerson("jason@raimondi.us")
	_, err = r.User().create(u)
	return err
}
