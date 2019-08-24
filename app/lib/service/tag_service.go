package service

import (
	"github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type PhotoAppService struct {
	repository *repository.Factory
}

func (s *PhotoAppService) AddAppsToPhoto(photoId string, tags []string) error {
	return s.linkToPhoto("app", photoId, tags)
}
func (s *PhotoAppService) RemoveAppFromPhoto(photoId string, appId int64) error {
	return s.repository.AppRepository().UnlinkFromPhoto(photoId, appId)
}

func (s *PhotoAppService) AddTagsToPhoto(photoId string, tags []string) error {
	return s.linkToPhoto("tag", photoId, tags)
}

func (s *PhotoAppService) RemoveTagFromPhoto(photoId string, tagId uint) error {
	return s.repository.TagRepository().UnlinkFromPhoto(photoId, tagId)
}

func (s *PhotoAppService) linkToPhoto(table string, photoId string, relate []string) error {
	queryBuilder := s.repository.QueryBuilder()
	var photo models.Photo
	s.repository.DB().First(&photo, "id = ?", uuid.FromStringOrNil(photoId))
	if newNames, err := s.getTagNamesToCreate(table, queryBuilder, relate); err != nil {
		return err
	} else {
		s.createNameRecords(table, queryBuilder, newNames)
	}

	if namesToLink, err := s.existingPhotoTag(table, queryBuilder, relate, photoId); err != nil {
		return err
	} else if len(namesToLink) > 0 {
		if tagsToLink, err := s.getIdsToLink(table, queryBuilder, namesToLink); err != nil {
			return err
		} else {
			for _, tag := range tagsToLink {
				photo.AddTag(tag)
			}
			s.repository.DB().Save(photo)
		}
		//} else if err = s.createLinkedRecords(table, queryBuilder, photoId, tagsToLink); err != nil {
		//	return err
		//}
	}

	return nil
}

func (s *PhotoAppService) createLinkedRecords(table string, q squirrel.StatementBuilderType, photoId string, idsToLink []models.Tag) error {
	insert := q.Insert("photo_"+table).Columns("photo_id", table+"_id")
	for _, id := range idsToLink {
		uid := uuid.FromStringOrNil(photoId)
		insert = insert.Values(uid, id.ID)
	}
	sql, args, err := insert.ToSql()
	if err != nil {
		return err
	}
	return s.repository.DB().Raw(sql, args...).Error
}

type Result struct {
	Id int64
}

func (s *PhotoAppService) getIdsToLink(table string, q squirrel.StatementBuilderType, names []string) ([]models.Tag, error) {
	tagsToLink := []models.Tag{}
	err := s.repository.DB().Find(&tagsToLink, "name IN (?)", names).Error
	return tagsToLink, err
}

func (s *PhotoAppService) createNameRecords(table string, q squirrel.StatementBuilderType, names []string) {
	for _, name := range names {
		s.repository.DB().Create(&models.Tag{Name: name})
	}
}

type existingNames struct {
	Name string
}

func (s *PhotoAppService) getTagNamesToCreate(table string, q squirrel.StatementBuilderType, names []string) (newTags []string, err error) {
	var tags []models.Tag
	if err = s.repository.DB().Where("name IN (?)", names).Find(&tags).Error; err != nil {
		return nil, err
	}
	var existingTagString []string
	for _, t := range tags {
		existingTagString = append(existingTagString, t.Name)
	}
	newTags = Difference(names, existingTagString)
	return newTags, err
}

func (s *PhotoAppService) existingPhotoTag(table string, q squirrel.StatementBuilderType, tags []string, photoId string) ([]string, error) {
	sql, args, err := q.Select("name").From("tags").LeftJoin("photo_tag on photo_tag.tag_id=tags.id").Where(squirrel.Eq{
		"photo_tag.photo_id": photoId,
	}).ToSql()
	if err != nil {
		return nil, err
	}
	existingLinkedTags := []existingNames{}
	if err = s.repository.DB().Raw(sql, args...).Scan(&existingLinkedTags).Error; err != nil {
		return nil, err
	}
	var existingTagString []string
	for _, t := range existingLinkedTags {
		existingTagString = append(existingTagString, t.Name)
	}
	existingTags := Difference(tags, existingTagString)
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
