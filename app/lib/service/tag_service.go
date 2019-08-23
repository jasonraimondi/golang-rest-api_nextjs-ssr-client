package service

import (
	"github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
)

type PhotoAppService struct {
	repository *repository.Factory
}

//type PhotoTagService struct {
//	repository *repository.Factory
//}

func (s *PhotoAppService) AddAppsToPhoto(photoId string, tags []string) error {
	return s.linkToPhoto("app", photoId, tags)
}
func (s *PhotoAppService) RemoveAppFromPhoto(photoId string, appId int64) error {
	return s.repository.AppRepository().UnlinkFromPhoto(photoId, appId)
}

func (s *PhotoAppService) AddTagsToPhoto(photoId string, tags []string) error {
	return s.linkToPhoto("tag", photoId, tags)
}

func (s *PhotoAppService) RemoveTagFromPhoto(photoId string, tagId int64) error {
	return s.repository.TagRepository().UnlinkFromPhoto(photoId, tagId)
}




func (s *PhotoAppService) linkToPhoto(table string, photoId string, relate []string) error {
	queryBuilder := s.repository.QueryBuilder()
	if newNames, err := s.getNamesToCreate(table, queryBuilder, relate); err != nil {
		return err
	} else if err = s.createNameRecords(table, queryBuilder, newNames); err != nil {
		return err
	}

	if namesToLink, err := s.existingPhotoTag(table, queryBuilder, relate, photoId); err != nil {
		return err
	} else if len(namesToLink) > 0 {
		if idsToLink, err := s.getIdsToLink(table, queryBuilder, namesToLink); err != nil {
			return err
		} else if	err = s.createLinkedRecords(table, queryBuilder, photoId, idsToLink); err != nil {
			return err
		}
	}

	return nil
}

func (s *PhotoAppService) createLinkedRecords(table string, q squirrel.StatementBuilderType, photoId string, idsToLink []int64) error {
	insert := q.Insert("photo_"+table).Columns("photo_id", table+"_id")
	for _, id := range idsToLink {
		insert = insert.Values(uuid.FromStringOrNil(photoId), id)
	}
	sql, args, err := insert.ToSql()
	if err != nil {
		return err
	}
	s.repository.DB().MustExec(sql, args...)
	return nil
}

func (s *PhotoAppService) getIdsToLink(table string, q squirrel.StatementBuilderType, namesToLink []string) ([]int64, error) {
	idsToLink := []int64{}
	sql, args, err := q.Select("id").From(table+"s").Where(squirrel.Eq{
		"name": namesToLink,
	}).ToSql()
	if err != nil {
		return nil, err
	}
	err = s.repository.DB().Select(&idsToLink, sql, args...)
	return idsToLink, err
}

func (s *PhotoAppService) createNameRecords(table string, q squirrel.StatementBuilderType, tags []string) error {
	insert := q.Insert(table+"s").Columns("name")
	for _, tag := range tags {
		insert = insert.Values(tag)
	}
	sql, args, _ := insert.ToSql()
	_ = s.repository.DB().MustExec(sql, args...)
	return nil
}

func (s *PhotoAppService) getNamesToCreate(table string, q squirrel.StatementBuilderType, tags []string) ([]string, error) {
	sql, args, err := q.Select("name").From(table+"s").Where(squirrel.Eq{
		"name": tags,
	}).ToSql()
	if err != nil {
		return nil, err
	}
	existingTags := &[]string{}
	err = s.repository.DB().Select(existingTags, sql, args...)
	newTags := Difference(tags, *existingTags)
	return newTags, err
}

func (s *PhotoAppService) existingPhotoTag(table string, q squirrel.StatementBuilderType, tags []string, photoId string) ([]string, error) {
	sql, args, err := q.Select("name").From(table+"s").LeftJoin("photo_"+table+" on photo_"+table+"."+table+"_id="+table+"s.id").Where(squirrel.Eq{
		"photo_"+table+".photo_id": photoId,
	}).ToSql()
	if err != nil {
		return nil, err
	}
	existingLinkedTags := &[]string{}
	err = s.repository.DB().Select(existingLinkedTags, sql, args...)
	existingTags := Difference(tags, *existingLinkedTags)
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
