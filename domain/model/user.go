package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/satori/go.uuid"
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

func (u *User) GetID() string {
	return u.ID.String()
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
