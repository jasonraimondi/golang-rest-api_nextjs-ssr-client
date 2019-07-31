package models

import (
	"database/sql"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

type User struct {
	ID           uuid.UUID      `db:"id"`
	First        sql.NullString `db:"first"`
	Last         sql.NullString `db:"last"`
	Email        string         `db:"email"`
	PasswordHash sql.NullString `db:"password_hash"`
	IsVerified   bool           `db:"is_verified"`
	CreatedAt    int64          `db:"created_at"`
	ModifiedAt   sql.NullInt64  `db:"modified_at"`
}

func NewUser(email string) (u *User) {
	return &User{
		ID:           uuid.NewV4(),
		First:        ToNullString(""),
		Last:         ToNullString(""),
		Email:        strings.ToLower(email),
		PasswordHash: ToNullString(""),
		IsVerified:   false,
		CreatedAt:    time.Now().Unix(),
		ModifiedAt:   ToNullInt64(""),
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
