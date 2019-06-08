package repository

import (
	"database/sql"
	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type Application struct {
	dbx    *sqlx.DB
	driver database.Driver
}

func NewTestDB() (r *RepositoryFactory, err error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	} else if err = db.Ping(); err != nil {
		panic(err)
	}
	dbx := sqlx.NewDb(db, "sqlite3")
	driver, err := sqlite3.WithInstance(dbx.DB, &sqlite3.Config{})
	if err != nil {
		return nil, err
	}
	a := NewApplication(dbx, driver)
	if err = a.MigrateNow(); err != nil {
		return nil, err
	}
	return r, err
}

func NewApplication(dbx *sqlx.DB, driver database.Driver) *Application {
	return &Application{dbx: dbx, driver: driver}
}

func (a *Application) MigrateNow() error {
	m, err := a.Migrate(a.driver)
	if err != nil {
		return err
	}
	return m.Up()
}

func (a *Application) Migrate(driver database.Driver) (*migrate.Migrate, error) {
	return migrate.NewWithDatabaseInstance("file://../migrations", "ql", driver)
}

