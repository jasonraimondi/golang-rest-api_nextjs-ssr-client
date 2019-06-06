package web

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
}

type UserRepository struct {
	db *sqlx.DB
}

func (r UserRepository) create(u User) {
	r.db.MustExec()
}