package models

import (
	"database/sql"
)

type Person struct {
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	Email     string
}

func NewSimplePerson(email string) (p Person) {
	return Person{
		Email:     email,
		FirstName: ToNullString(""),
		LastName:  ToNullString(""),
	}
}