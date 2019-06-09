package db

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
	return migrate.NewWithDatabaseInstance("file://../db/migrations", "ql", databaseInstance)
}
