package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type SignUpConfirmation struct {
	Token     uuid.UUID `db:"token"`
	CreatedAt int64     `db:"created_at"`
	UserId    uuid.UUID `db:"user_id"`
}

func NewSignUpConfirmation(u User) (c *SignUpConfirmation) {
	return &SignUpConfirmation{
		Token:     uuid.NewV4(),
		CreatedAt: time.Now().Unix(),
		UserId:    u.ID,
	}
}
