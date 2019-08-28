package repository

import (
	"github.com/jinzhu/gorm"
)

type Factory struct {
	debug bool
	db    *gorm.DB
}

func NewFactory(dbx *gorm.DB, debug bool) *Factory {
	return &Factory{debug, dbx}
}

func (r *Factory) DB() *gorm.DB {
	return r.db
}

func (r *Factory) UserRepository() *UserRepository {
	return &UserRepository{r.db}
}

func (r *Factory) SignUpConfirmation() *SignUpConfirmationRepository {
	return &SignUpConfirmationRepository{r.db}
}

func (r *Factory) PhotoRepository() *PhotoRepository {
	return &PhotoRepository{r.debug, r.db}
}
