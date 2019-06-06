package web_test

import (
	"database/sql"
	"git.jasonraimondi.com/jason/learn-with-tests/web"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

var schema = `
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)`

func TestGetConnection(t *testing.T) {
	driver, _ := sql.Open("sqlite3", ":memory:")
	db := sqlx.NewDb(driver, "sqlite3")
	r := web.RepositoryFactory{DB: db}
	err := r.Migrate(schema)
	err = r.Seed()
	u := r.User()
	assert.Nil(t, err)
	assert.True(t, u != nil)
	assert.True(t, true)
}
