package repository

import (
	"github.com/jinzhu/gorm"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type AppRepository struct {
	debug bool
	db    *gorm.DB
}

func (r *AppRepository) GetById(id string) (photo *models.Photo, err error) {
	photo = &models.Photo{}
	err = r.db.First(&photo, "id = ?", id).Error
	return photo, err
}

func (r *AppRepository) GetByName(name string) (photo *models.Photo, err error) {
	photo = &models.Photo{}
	err = r.db.First(&photo, "name = ?", name).Error
	return photo, err
}
