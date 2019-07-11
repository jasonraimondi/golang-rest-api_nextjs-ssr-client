package lib

import (
	"database/sql"
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

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
