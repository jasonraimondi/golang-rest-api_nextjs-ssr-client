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
	Email        string         `db:"email"`
	CreatedAt    int64          `db:"created_at"`
	ModifiedAt   sql.NullInt64  `db:"modified_at"`
}

func NewPerson(email string) (p Person) {
	return Person{
		Id:           uuid.NewV4().String(),
		FirstName:    ToNullString(""),
		LastName:     ToNullString(""),
		Email:        email,
		CreatedAt:    time.Now().Unix(),
		ModifiedAt:   ToNullInt64(""),
	}
}
