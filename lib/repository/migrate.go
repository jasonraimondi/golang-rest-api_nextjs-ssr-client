package repository

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
)

func MigrateNow(driver *database.Driver, dir string) error {
	m, err := Migrate(*driver, dir)
	if err != nil {
		return err
	}
	return m.Up()
}

func Migrate(databaseInstance database.Driver, dir string) (*migrate.Migrate, error) {
	dir = "/Users/jason/go/src/git.jasonraimondi.com/jason/jasontest/db/migrations"
	return migrate.NewWithDatabaseInstance("file://" + dir, "ql", databaseInstance)
}
