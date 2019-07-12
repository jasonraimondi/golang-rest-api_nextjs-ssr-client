package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type SignUpConfirmation struct {
	Token     uuid.UUID `db:"token"`
	CreatedAt time.Time `db:"created_at"`
	*User     `db:"user_id"`
}

func NewSignUpConfirmation(u *User) (c *SignUpConfirmation) {
	return &SignUpConfirmation{
		Token:     uuid.NewV4(),
		CreatedAt: time.Now(),
		User:      u,
	}
}
