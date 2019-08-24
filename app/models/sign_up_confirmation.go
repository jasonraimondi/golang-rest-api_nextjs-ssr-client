package models

import (
	uuid "github.com/satori/go.uuid"
)

type SignUpConfirmation struct {
	Token  uuid.UUID `gorm:"primary_key"`
	UserID uuid.UUID
}

func NewSignUpConfirmation(u User) (c *SignUpConfirmation) {
	return &SignUpConfirmation{
		Token:  uuid.NewV4(),
		UserID: u.GetID(),
	}
}
