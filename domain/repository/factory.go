package repository

import (
	"github.com/jmoiron/sqlx"
)

type Factory struct {
	DBx *sqlx.DB
}

func NewFactory(dbx *sqlx.DB) *Factory {
	return &Factory{dbx}
}

func (r *Factory) User() *SqlxUserRepository {
	return NewSqlxUserRepository(r.DBx)
}

func (r *Factory) SignUpConfirmation() *SqlxSignUpConfirmationRepository {
	return NewSqlxSignUpConfirmationRepository(r.DBx)
}
