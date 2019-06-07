package repository_test

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository(t *testing.T) {
	_, err := ConnectToSQL()
	assert.Nil(t, err)
}

func ConnectToSQL() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance("file://../migrations", "ql", driver)
	if err != nil {
		return nil, err
	}
	if err = m.Up(); err != nil {
		return nil, err
	}
	return db, err
}
