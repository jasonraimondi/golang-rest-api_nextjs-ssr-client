package models

import (
	"database/sql"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Person struct {
	Id           string         `db:"id"`
	FirstName    sql.NullString `db:"first_name"`
	LastName     sql.NullString `db:"last_name"`
	PasswordHash sql.NullString `db:"password_hash"`
	IsVerified   bool           `db:"is_verified"`
	Email        string         `db:"email"`
	CreatedAt    int64          `db:"created_at"`
	ModifiedAt   sql.NullInt64  `db:"modified_at"`
}

func NewPerson(email string) (p Person) {
	return Person{
		Id:         uuid.NewV4().String(),
		FirstName:  ToNullString(""),
		LastName:   ToNullString(""),
		Email:      email,
		CreatedAt:  time.Now().Unix(),
		ModifiedAt: ToNullInt64(""),
	}
}

func (p *Person) SetPassword(pass string) (err error) {
	bytes, err := HashPassword(pass)
	s := string(bytes)
	p.PasswordHash = ToNullString(s)
	return err
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
