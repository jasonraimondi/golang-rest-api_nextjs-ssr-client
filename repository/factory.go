package repository

import (
	"github.com/jmoiron/sqlx"
)

type Factory struct {
	dbx *sqlx.DB
}

func NewFactory(dbx *sqlx.DB) *Factory {
	return &Factory{dbx}
}

func (r Factory) Person() *Person {
	return &Person{dbx: r.dbx}
}
