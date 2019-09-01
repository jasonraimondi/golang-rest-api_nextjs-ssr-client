package service

import (
	"github.com/jinzhu/gorm"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type TagService struct {
	db              *gorm.DB
	photoRepository *repository.PhotoRepository
}

func (s *TagService) AddTagsToPhoto(photo *models.Photo, tags []string) error {
	if err := s.createMissingTags(tags); err != nil {
		return err
	}
	tagsToLink, err := s.GetAllNewTagNamesToCreate(tags, photo.GetID())
	if err != nil {
		return err
	}
	if len(tagsToLink) > 0 {
		tagsToLink, err := s.GetAllTagsByName(tagsToLink)
		if err != nil {
			return err
		}
		photo.AddTags(tagsToLink)
		s.db.Save(photo)
	}
	return nil
}

func (s *TagService) createMissingTags(tags []string) error {
	newNames, err := s.GetAllTagNamesToCreate(tags)
	if err != nil {
		return err
	}
	s.CreateTagsForNames(newNames)
	return nil
}

func (s *TagService) RemoveTagFromPhoto(photoId string, tagId uint) error {
	return s.photoRepository.UnlinkTag(photoId, tagId)
}

func (s *TagService) createLinkedRecords(photoId string, tagsToLink []models.Tag) error {
	return s.db.Association("tags").Append(tagsToLink).Error
}

func (s *TagService) GetAllTagsByName(names []string) ([]models.Tag, error) {
	var tagsToLink []models.Tag
	err := s.db.Find(&tagsToLink, "name IN (?)", names).Error
	return tagsToLink, err
}

func (s *TagService) CreateTagsForNames(names []string) {
	for _, name := range names {
		s.db.Create(&models.Tag{Name: name})
	}
}

func (s *TagService) GetAllTagNamesToCreate(names []string) (result []string, err error) {
	var tags []models.Tag
	if err = s.db.Where("name IN (?)", names).Find(&tags).Error; err != nil {
		return nil, err
	}
	var existingTagString []string
	for _, t := range tags {
		existingTagString = append(existingTagString, t.Name)
	}
	result = ArrayDiff(names, existingTagString)
	return result, err
}

func (s *TagService) GetAllNewAppNamesToCreate(name []string, photoId string) (result []string, err error) {
	var existingAppString []string
	err = s.db.
		Model(&models.PhotoApp{}).
		Joins("left join tags on tags.id=photo_tag.tag_id").
		Where("photo_app.photo_id = ?", photoId).
		Pluck("tags.name", &existingAppString).
		Error
	if err != nil {
		return nil, err
	}
	result = ArrayDiff(name, existingAppString)
	if result == nil {
		result = []string{}
	}
	return result, err
}

func (s *TagService) GetAllNewTagNamesToCreate(name []string, photoId string) (result []string, err error) {
	var existingTagString []string
	err = s.db.
		Model(&models.PhotoTag{}).
		Joins("left join tags on tags.id=photo_tag.tag_id").
		Where("photo_tag.photo_id = ?", photoId).
		Pluck("tags.name", &existingTagString).
		Error
	if err != nil {
		return nil, err
	}
	result = ArrayDiff(name, existingTagString)
	if result == nil {
		result = []string{}
	}
	return result, err
}

// difference returns the elements in `a` that aren't in `b`.
func ArrayDiff(a, b []string) (diff []string) {
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
