package models

import (
	uuid "github.com/satori/go.uuid"
)

type SignUpConfirmation struct {
	Token  uuid.UUID
	UserID uuid.UUID
	User   *User
}


func NewSignUpConfirmation(u *User) (c *SignUpConfirmation) {
	return &SignUpConfirmation{
		Token: uuid.NewV4(),
		User:  u,
	}
}
