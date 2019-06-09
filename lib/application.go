package lib

import (
	"database/sql"
	"fmt"
	"git.jasonraimondi.com/jason/jasontest/repository"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/jmoiron/sqlx"
	"os"
	"path/filepath"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type Application struct {
	dbx    *sqlx.DB
	driver *database.Driver
}

func (a *Application) RepositoryFactory() *repository.Factory {
	return repository.NewFactory(a.dbx)
}

func (a *Application) MigrateNow() error {
	m, err := a.Migrate(*a.driver)
	if err != nil {
		return err
	}
	return m.Up()
}

func (a *Application) Migrate(databaseInstance database.Driver) (*migrate.Migrate, error) {
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	return migrate.NewWithDatabaseInstance("file://../migrations", "ql", databaseInstance)
}

func NewApplication(dbx *sqlx.DB, driver *database.Driver) *Application {
	return &Application{dbx: dbx, driver: driver}
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
	a = NewApplication(dbx, &databaseInstance)
	if err = a.MigrateNow(); err != nil {
		panic(err)
	}
	return a
}
