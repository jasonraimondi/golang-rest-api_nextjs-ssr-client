package service

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type PhotoAppService struct {
	db              *gorm.DB
	photoRepository *repository.PhotoRepository
}

func (s *PhotoAppService) AddTagsToPhoto(photoId string, tags []string) error {
	var photo models.Photo
	s.db.First(&photo, "id = ?", uuid.FromStringOrNil(photoId))
	newNames, err := s.getTagNamesToCreate(tags)
	if err != nil {
		return err
	}
	s.createNameRecords(newNames)
	tagsToLink, err := s.existingPhotoTag(tags, photoId)
	if err != nil {
		return err
	}
	if len(tagsToLink) > 0 {
		tagsToLink, err := s.getIdsToLink(tagsToLink)
		if err != nil {
			return err
		}
		photo.AddTags(tagsToLink)
		s.db.Save(photo)
	}
	return nil
}

func (s *PhotoAppService) RemoveTagFromPhoto(photoId string, tagId uint) error {
	return s.photoRepository.UnlinkFromPhoto(photoId, tagId)
}

func (s *PhotoAppService) createLinkedRecords(photoId string, tagsToLink []models.Tag) error {
	return s.db.Association("tags").Append(tagsToLink).Error
}

func (s *PhotoAppService) getIdsToLink(names []string) ([]models.Tag, error) {
	tagsToLink := []models.Tag{}
	err := s.db.Find(&tagsToLink, "name IN (?)", names).Error
	return tagsToLink, err
}

func (s *PhotoAppService) createNameRecords(names []string) {
	for _, name := range names {
		s.db.Create(&models.Tag{Name: name})
	}
}

func (s *PhotoAppService) getTagNamesToCreate(names []string) (result []string, err error) {
	var tags []models.Tag
	if err = s.db.Where("name IN (?)", names).Find(&tags).Error; err != nil {
		return nil, err
	}
	var existingTagString []string
	for _, t := range tags {
		existingTagString = append(existingTagString, t.Name)
	}
	result = Difference(names, existingTagString)
	return result, err
}

func (s *PhotoAppService) existingPhotoTag(name []string, photoId string) ([]string, error) {
	var existingTagString []string
	err := s.db.Model(&models.Tag{}).Joins("left join photo_tag on photo_tag.tag_id=tags.id").Where("photo_tag.photo_id = ?", photoId).Pluck("tags.name", &existingTagString).Error
	if err != nil {
		return []string{}, err
	}
	existingTags := Difference(name, existingTagString)
	if existingTags == nil {
		existingTags = []string{}
	}
	return existingTags, err
}

// difference returns the elements in `a` that aren't in `b`.
func Difference(a, b []string) (diff []string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}
