package service

import (
	"github.com/jinzhu/gorm"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type AppService struct {
	db              *gorm.DB
	photoRepository *repository.PhotoRepository
}

func (s *AppService) AddAppToPhoto(photo *models.Photo, name string) error {

	return s.photoRepository.Update(photo)
}
