package repository

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func Initialize() (r *RepositoryFactory, err error) {
	driver, err := ConnectToSQL()
	if err != nil {
		return nil, err
	}
	r = &RepositoryFactory{DB: driver}
	if err := MigrateNow(r.DB.DB); err != nil {
		return nil, err
	}
	return r, nil
}

func ConnectToSQL() (driver *sqlx.DB, err error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	driver = sqlx.NewDb(db, "sqlite3")
	return driver, nil
}

func MigrateNow(db *sql.DB) error {
	m, err := Migrate(db)
	if err != nil {
		return err
	}
	if err = m.Up(); err != nil {
		return err
	}
	return nil
}

func Migrate(db *sql.DB) (*migrate.Migrate, error) {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, err
	}
	return migrate.NewWithDatabaseInstance("file://../migrations", "ql", driver)
}
