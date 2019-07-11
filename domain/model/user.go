package model

import (
	"database/sql"
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID           string         `db:"id"`
	FirstName    sql.NullString `db:"first_name"`
	LastName     sql.NullString `db:"last_name"`
	Email        string         `db:"email"`
	PasswordHash sql.NullString `db:"password_hash"`
	IsVerified   bool           `db:"is_verified"`
	CreatedAt    int64          `db:"created_at"`
	ModifiedAt   sql.NullInt64  `db:"modified_at"`
}

func NewUser(email string) (u User) {
	return User{
		ID:           uuid.NewV4().String(),
		FirstName:    ToNullString(""),
		LastName:     ToNullString(""),
		Email:        email,
		PasswordHash: ToNullString(""),
		CreatedAt:    time.Now().Unix(),
		ModifiedAt:   ToNullInt64(""),
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
