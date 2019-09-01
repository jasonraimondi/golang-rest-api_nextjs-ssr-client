package models

import (
	uuid "github.com/satori/go.uuid"
)

type SignUpConfirmation struct {
	Token  uuid.UUID `gorm:"primary_key"`
	UserID uuid.UUID `gorm:"index:user_idx; unique; not null"`
	User   *User
}

func NewSignUpConfirmation(u *User) (c *SignUpConfirmation) {
	return &SignUpConfirmation{
		Token: uuid.NewV4(),
		User:  u,
	}
}
