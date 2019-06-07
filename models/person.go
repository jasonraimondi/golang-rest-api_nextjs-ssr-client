package models

import (
	"database/sql"
	"github.com/satori/go.uuid"
	"time"
)

type Person struct {
	Id           string         `db:"id"`
	FirstName    sql.NullString `db:"first_name"`
	LastName     sql.NullString `db:"last_name"`
	PasswordHash sql.NullString `db:"password_hash"`
	Email        string         `db:"email"`
	CreatedAt    int64          `db:"created_at"`
	ModifiedAt   sql.NullInt64  `db:"modified_at"`
}

func NewPerson(email string) (p Person) {
	return Person{
		Id:        uuid.NewV4().String(),
		Email:     email,
		FirstName: ToNullString(""),
		LastName:  ToNullString(""),
		CreatedAt: time.Now().Unix(),
		ModifiedAt: ToNullInt64(""),
	}
}
