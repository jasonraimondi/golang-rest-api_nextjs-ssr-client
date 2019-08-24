package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/gorm"
)

type TagRepository struct {
	queryBuilder squirrel.StatementBuilderType
	dbx          *gorm.DB
}

func (r *TagRepository) Delete(id string) error {
	sql, args, err := r.queryBuilder.Delete("tags").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}
	return r.dbx.Raw(sql, args...).Error
}

func (r *TagRepository) UnlinkFromPhoto(photoId string, tagId int64) error {
	sql, args, err := r.queryBuilder.
		Delete("photo_tag").
		Where(squirrel.Eq{
			"tag_id":   tagId,
			"photo_id": photoId,
		}).ToSql()
	if err != nil {
		return err
	}
	return r.dbx.Raw(sql, args...).Error
}
