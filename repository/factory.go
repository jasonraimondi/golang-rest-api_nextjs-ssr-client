package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RepositoryFactory struct {
	dbx *sqlx.DB
}

func (r RepositoryFactory) Person() *PersonRepository {
	fmt.Println("hello jason")
	return &PersonRepository{dbx: r.dbx}
}
