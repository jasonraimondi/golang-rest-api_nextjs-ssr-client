package models

import (
	"database/sql"
	"github.com/satori/go.uuid"
)

type Person struct {
	ID        string         `db:"id"`
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	Email     string         `db:"email"`
}

func NewSimplePerson(email string) (p Person) {
	return Person{
		ID:        uuid.NewV4().String(),
		Email:     email,
		FirstName: ToNullString(""),
		LastName:  ToNullString(""),
	}
}
