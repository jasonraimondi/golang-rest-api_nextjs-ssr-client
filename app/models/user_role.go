package models

import (
	uuid "github.com/satori/go.uuid"
)

type UserRole struct {
	UserID uuid.UUID
	User   *User
	RoleID uint
	Role   *Role
}

func (p *UserRole) TableName() string {
	return "user_role"
}
