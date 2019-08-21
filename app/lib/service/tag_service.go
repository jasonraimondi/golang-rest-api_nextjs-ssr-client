package service

import (
	"fmt"

	"github.com/Masterminds/squirrel"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
)

type TagService struct {
	repository *repository.Factory
}

func (s *TagService) AddTagsToPhoto(desiredTags []string) error {
	photoId := "08f62e0b-58c3-46c1-b1f8-1d616f17931a"

	query := s.repository.QueryBuilder().Select("name").From("tags").LeftJoin("photo_tag on photo_tag.tag_id=tags.id ").Where(squirrel.Eq{
		"photo_tag.photo_id": photoId,
	})
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	existingTags := &[]string{}
	err = s.repository.DB().Select(existingTags, sql, args...)
	newTags := Difference(desiredTags, *existingTags)
	if err != nil {
		return err
	}
	// @todo loop through and create these tags
	fmt.Println(newTags)
	return nil
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