package service

import (
	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type PhotoAppService struct {
	repository *repository.Factory
}

func (s *PhotoAppService) AddTagsToPhoto(photoId string, tags []string) error {
	return s.linkToPhoto("tag", photoId, tags)
}

func (s *PhotoAppService) RemoveTagFromPhoto(photoId string, tagId uint) error {
	return s.repository.TagRepository().UnlinkFromPhoto(photoId, tagId)
}

func (s *PhotoAppService) linkToPhoto(table string, photoId string, names []string) error {
	var photo models.Photo
	s.repository.DB().First(&photo, "id = ?", uuid.FromStringOrNil(photoId))
	newNames, err := s.getTagNamesToCreate(names)
	if err != nil {
		return err
	}
	s.createNameRecords(newNames)
	tagsToLink, err := s.existingPhotoTag(names, photoId)
	if err != nil {
		return err
	}
	if len(tagsToLink) > 0 {
		tagsToLink, err := s.getIdsToLink(tagsToLink)
		if err != nil {
			return err
		}
		photo.AddTags(tagsToLink)
		s.repository.DB().Save(photo)
	}
	return nil
}

func (s *PhotoAppService) createLinkedRecords(photoId string, tagsToLink []models.Tag) error {
	return s.repository.DB().Association("tags").Append(tagsToLink).Error
}

type Result struct {
	Id int64
}

func (s *PhotoAppService) getIdsToLink(names []string) ([]models.Tag, error) {
	tagsToLink := []models.Tag{}
	err := s.repository.DB().Find(&tagsToLink, "name IN (?)", names).Error
	return tagsToLink, err
}

func (s *PhotoAppService) createNameRecords(names []string) {
	for _, name := range names {
		s.repository.DB().Create(&models.Tag{Name: name})
	}
}

func (s *PhotoAppService) getTagNamesToCreate(names []string) (result []string, err error) {
	var tags []models.Tag
	if err = s.repository.DB().Where("name IN (?)", names).Find(&tags).Error; err != nil {
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
	err := s.repository.DB().Model(&models.Tag{}).Joins("left join photo_tag on photo_tag.tag_id=tags.id").Where("photo_tag.photo_id = ?", photoId).Pluck("tags.name", &existingTagString).Error
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
