package service

import (
	"github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
)

type TagService struct {
	repository *repository.Factory
}

func (s *TagService) AddTagsToPhoto(photoId string, tags []string) error {
	queryBuilder := s.repository.QueryBuilder()
	if newTags, err := s.findTagsToCreate(queryBuilder, tags); err != nil {
		return err
	} else if err = s.createTags(queryBuilder, newTags); err != nil {
		return err
	}

	tagNamesToLink, err := s.existingPhotoTag(queryBuilder, tags, photoId)
	if err != nil {
		return err
	}

	tagIdsToLink := &[]int64{}
	if len(tagNamesToLink) > 0 {
		// GET TAGS IDS OF TO LINK
		sql, args, err := queryBuilder.Select("id").From("tags").Where(squirrel.Eq{
			"name": tagNamesToLink,
		}).ToSql()
		if err != nil {
			return err
		}
		err = s.repository.DB().Select(tagIdsToLink, sql, args...)


		// ADD LINKED RECORDS
		q := queryBuilder.Insert("photo_tag").Columns("photo_id", "tag_id")
		for _, id := range *tagIdsToLink {
			q = q.Values(uuid.FromStringOrNil(photoId), id)
		}
		sql, args, err = q.ToSql()
		if err != nil {
			return err
		}
		s.repository.DB().MustExec(sql, args...)
	}

	return nil
}

func (s *TagService) createTags(q squirrel.StatementBuilderType, tags []string) error {
	insertQuery := q.Insert("tags").Columns("name")
	for _, tag := range tags {
		insertQuery = insertQuery.Values(tag)
	}
	sql, args, _ := insertQuery.ToSql()
	_ = s.repository.DB().MustExec(sql, args...)
	return nil
}

func (s *TagService) findTagsToCreate(q squirrel.StatementBuilderType, tags []string) ([]string, error) {
	sql, args, err := q.Select("name").From("tags").Where(squirrel.Eq{
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

func (s *TagService) existingPhotoTag(q squirrel.StatementBuilderType, tags []string, photoId string) ([]string, error) {
	sql, args, err := q.Select("name").From("tags").LeftJoin("photo_tag on photo_tag.tag_id=tags.id ").Where(squirrel.Eq{
		"photo_tag.photo_id": photoId,
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