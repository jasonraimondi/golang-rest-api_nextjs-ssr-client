package repository

import (
	"github.com/jinzhu/gorm"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type PhotoRepository struct {
	dbx *gorm.DB
}

func (r *PhotoRepository) GetById(id string) (photo models.Photo, err error) {
	photo = models.Photo{}
	err = r.dbx.First(photo).Error
	return photo, err
}

func (r *PhotoRepository) Update(u *models.Photo) (err error) {
	return r.dbx.Update(u).Error
}

func (r *PhotoRepository) Create(u *models.Photo) (err error) {
	return r.dbx.Create(u).Error
}
