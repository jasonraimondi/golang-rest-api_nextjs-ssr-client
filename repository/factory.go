package repository

import (
	"github.com/jmoiron/sqlx"
)

type RepositoryFactory struct {
	dbx *sqlx.DB
}

func (r RepositoryFactory) Person() *PersonRepository {
	return &PersonRepository{dbx: r.dbx}
}
