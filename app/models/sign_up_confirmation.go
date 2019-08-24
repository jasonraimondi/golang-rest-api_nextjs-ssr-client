package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SignUpConfirmation struct {
	gorm.Model
	Token     uuid.UUID
	UserID    uuid.UUID
}

func NewSignUpConfirmation(u User) (c *SignUpConfirmation) {
	return &SignUpConfirmation{
		Token:     uuid.NewV4(),
		UserID:    u.GetID(),
	}
}
