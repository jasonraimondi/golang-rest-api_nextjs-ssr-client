package repository_test

import (
	"database/sql"
	"git.jasonraimondi.com/jason/jasontest/models"
	"git.jasonraimondi.com/jason/jasontest/repository"
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
	r := repository.RepositoryFactory{DB: db}
	err := r.Migrate(schema)
	err = r.Seed()
	u := r.User()





	assert.Nil(t, err)
	assert.True(t, u != nil)
	assert.True(t, true)




	// You can also get a single result, a la QueryRow
	jason := models.Person{}
	err = db.Get(&jason, "SELECT * FROM person WHERE first_name=$1", "Jason")

	// if you have null fields and use SELECT *, you must use sql.Null* in your struct
	places := []models.Place{}
	err = db.Select(&places, "SELECT * FROM place ORDER BY telcode ASC")
	assert.Nil(t, err)
	assert.Equal(t, jason.FirstName, "Jason")
}
