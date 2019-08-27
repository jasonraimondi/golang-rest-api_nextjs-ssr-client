package repository

import (
	"github.com/jinzhu/gorm"
)

type Factory struct {
	db *gorm.DB
}

func NewFactory(dbx *gorm.DB) *Factory {
	return &Factory{dbx}
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
	return &PhotoRepository{r.db}
}

func (r *Factory) TagRepository() *TagRepository {
	return &TagRepository{r.db}
}
