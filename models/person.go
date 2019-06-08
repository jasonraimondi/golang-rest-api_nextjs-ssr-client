package models

import (
	"database/sql"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Person struct {
	ID           string         `db:"id"`
	FirstName    sql.NullString `db:"first_name"`
	LastName     sql.NullString `db:"last_name"`
	Email        string         `db:"email"`
	PasswordHash sql.NullString `db:"password_hash"`
	IsVerified   bool           `db:"is_verified"`
	CreatedAt    int64          `db:"created_at"`
	ModifiedAt   sql.NullInt64  `db:"modified_at"`
}

func NewPerson(email string) (p Person) {
	return Person{
		ID:           uuid.NewV4().String(),
		FirstName:    ToNullString(""),
		LastName:     ToNullString(""),
		Email:        email,
		PasswordHash: ToNullString(""),
		CreatedAt:    time.Now().Unix(),
		ModifiedAt:   ToNullInt64(""),
	}
}

func (p *Person) SetPassword(pass string) (err error) {
	bytes, err := HashPassword(pass)
	if err != nil {
		return err
	}
	p.PasswordHash = ToNullString(string(bytes))
	return nil
}

func (p *Person) CheckPassword(pass string) bool {
	return CheckPasswordHash(pass, p.PasswordHash.String)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
