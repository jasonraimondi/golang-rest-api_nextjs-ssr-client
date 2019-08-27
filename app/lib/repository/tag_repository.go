package repository

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/lib/pagination"
	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type TagRepository struct {
	db *gorm.DB
}

func (r *TagRepository) Delete(id string) error {
	var tag models.Tag
	r.db.First(&tag)
	return r.db.Delete(tag).Error
}

func (s *TagRepository) ForPhoto(photoId string, currentPage int64, itemsPerPage int64) *pagination.Paginator {
	var tags []models.Tag
	db := s.db.Joins("left join photo_tag on photo_tag.tag_id=tags.id").Where("photo_tag.photo_id = ?", uuid.FromStringOrNil(photoId))
	return pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    int(currentPage),
		Limit:   int(itemsPerPage),
		OrderBy: []string{"tags.name asc"},
		ShowSQL: true,
	}, &tags)
}
