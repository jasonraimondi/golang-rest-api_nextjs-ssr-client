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

func (r *Factory) User() *UserRepository {
	return &UserRepository{r.DBx}
}

func (r *Factory) SignUpConfirmation() *SignUpConfirmationRepository {
	return &SignUpConfirmationRepository{r.DBx}
}

func (r *Factory) PhotoRepository() *PhotoRepository {
	return &PhotoRepository{r.DBx}
}
