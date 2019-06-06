package web

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type RepositoryFactory struct {
	DB *sqlx.DB
}

func (r RepositoryFactory) User() *UserRepository {
	return &UserRepository{db: r.DB}
}
