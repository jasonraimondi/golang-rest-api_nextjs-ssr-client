package repository

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
)

func MigrateNow(driver *database.Driver) error {
	m, err := Migrate(*driver)
	if err != nil {
		return err
	}
	return m.Up()
}

func Migrate(databaseInstance database.Driver) (*migrate.Migrate, error) {
	return migrate.NewWithDatabaseInstance("file:///Users/jason/go/src/git.jasonraimondi.com/jason/jasontest/db/migrations", "ql", databaseInstance)
}
