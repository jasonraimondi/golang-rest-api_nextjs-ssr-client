package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"primary_key"`
	First        sql.NullString
	Last         sql.NullString
	Email        string `gorm:"size:255"`
	PasswordHash sql.NullString
	IsVerified   bool
	Photos       []Photo
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
}

func NewUser(email string) (u *User) {
	return &User{
		ID:           uuid.NewV4(),
		First:        ToNullString(""),
		Last:         ToNullString(""),
		Email:        strings.ToLower(email),
		PasswordHash: ToNullString(""),
		IsVerified:   false,
		CreatedAt:    time.Now(),
	}
}

func (u *User) GetID() uuid.UUID {
	return u.ID
}

func (u *User) GetFirst() sql.NullString {
	return u.First
}

func (u *User) GetLast() sql.NullString {
	return u.Last
}

func (u *User) SetEmail(email string) {
	u.Email = strings.ToLower(email)
}

func (u *User) SetFirst(first string) {
	u.First = ToNullString(first)
}

func (u *User) SetLast(last string) {
	u.Last = ToNullString(last)
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

func (u *User) SetVerified() {
	u.IsVerified = true
}

func (u *User) GetFullName() (name string) {
	var s []string
	if u.First.Valid {
		s = append(s, u.First.String)
	}
	if u.Last.Valid {
		s = append(s, u.Last.String)
	}
	return strings.Join(s, " ")
}

func (u *User) GetFullIdentifier() (name string) {
	return fmt.Sprintf("%s <%s>", u.GetFullName(), u.Email)
}
