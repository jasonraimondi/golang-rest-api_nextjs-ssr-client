package models

import (
	"database/sql"
)

type Person struct {
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	Email     string
}

func NewPerson(email string) (p Person) {
	return Person{
		Email: email,
	}
}

func NewNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}
