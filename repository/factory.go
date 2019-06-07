package repository

import (
	"github.com/jmoiron/sqlx"
)

type RepositoryFactory struct {
	DB *sqlx.DB
}

func (r RepositoryFactory) Role() *PersonRepository {
	return &PersonRepository{db: r.DB}
}

func (r RepositoryFactory) Person() *PersonRepository {
	return &PersonRepository{db: r.DB}
}
