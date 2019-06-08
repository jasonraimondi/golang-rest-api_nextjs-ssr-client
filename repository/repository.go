package repository

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type Application struct {
	dbx    *sqlx.DB
	driver *database.Driver
}

func (a *Application) RepositoryFactory() *RepositoryFactory {
	return &RepositoryFactory{
		dbx: a.dbx,
	}
}

func (a *Application) MigrateNow() error {
	m, err := a.Migrate(*a.driver)
	if err != nil {
		return err
	}
	return m.Up()
}

func (a *Application) Migrate(databaseInstance database.Driver) (*migrate.Migrate, error) {
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
