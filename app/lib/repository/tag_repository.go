package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type TagRepository struct {
	queryBuilder squirrel.StatementBuilderType
	dbx          *gorm.DB
}

func (r *TagRepository) Delete(id string) error {
	var tag models.Tag
	r.dbx.First(&tag)
	return r.dbx.Delete(tag).Error
}

func (r *TagRepository) UnlinkFromPhoto(photoId string, tagId uint) error {
	var tag models.Tag
	var photo models.Photo
	r.dbx.First(&tag, "id = ?", tagId)
	r.dbx.First(&photo, "id = ?", uuid.FromStringOrNil(photoId))
	return r.dbx.Model(&photo).Association("tags").Delete(tag).Error
}
