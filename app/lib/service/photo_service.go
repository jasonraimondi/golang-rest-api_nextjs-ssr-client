package service

import (
	"github.com/jinzhu/gorm"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type PhotoService struct {
	db              *gorm.DB
	photoRepository *repository.PhotoRepository
	appService      *AppService
	TagService      *TagService
}

func (s *PhotoService) UpdatePhoto(photoId string, description string, app string, tags []string) error {
	photo, err := s.photoRepository.GetById(photoId)
	if err != nil {
		return err
	}
	if app != "" {
		var a models.App
		if err := s.db.FirstOrCreate(&a, models.App{Name: app}).Error; err != nil {
			return err
		}
		photo.SetApp(&a)
	}
	if len(tags) > 0 {
		if err := s.TagService.AddTagsToPhoto(photo, tags); err != nil {
			return err
		}
	}
	if description != "" {
		photo.Description = models.ToNullString(description)
	}
	return s.photoRepository.Update(photo)
}
