package model

import (
	"database/sql"
	"github.com/lib/pq"
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID           uuid.UUID      `db:"id"`
	FirstName    sql.NullString `db:"first_name"`
	LastName     sql.NullString `db:"last_name"`
	Email        string         `db:"email"`
	PasswordHash sql.NullString `db:"password_hash"`
	IsVerified   bool           `db:"is_verified"`
	CreatedAt    time.Time      `db:"created_at"`
	ModifiedAt   pq.NullTime    `db:"modified_at"`
}

func NewUser(email string) (u *User) {
	return &User{
		ID:           uuid.NewV4(),
		FirstName:    ToNullString(""),
		LastName:     ToNullString(""),
		Email:        email,
		PasswordHash: ToNullString(""),
		IsVerified:   false,
		CreatedAt:    time.Now(),
		ModifiedAt:   ToNullNullTime(),
	}
}

func (u *User) SetPassword(pass string) (err error) {
	bytes, err := HashPassword(pass)
	if err != nil {
		return err
	}
	u.PasswordHash = ToNullString(string(bytes))
	return nil
}

func (u *User) CheckPassword(pass string) bool {
	return CheckPasswordHash(pass, u.PasswordHash.String)
}
