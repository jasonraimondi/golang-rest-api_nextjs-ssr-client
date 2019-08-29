package repository

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/lib/pagination"
	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type PhotoRepository struct {
	debug bool
	db    *gorm.DB
}

func (r *PhotoRepository) GetById(id string) (photo *models.Photo, err error) {
	photo = &models.Photo{}
	err = r.db.
		Preload("Tags").
		Preload("Apps").
		First(&photo, "id = ?", uuid.FromStringOrNil(id)).Error
	return photo, err
}

func (r *PhotoRepository) Update(u *models.Photo) (err error) {
	return r.db.Update(u).Error
}

func (r *PhotoRepository) Create(u *models.Photo) (err error) {
	return r.db.Create(u).Error
}

func (r *PhotoRepository) UnlinkTag(photoId string, tagId uint) error {
	var tag models.Tag
	var photo models.Photo
	r.db.First(&tag, "id = ?", tagId)
	r.db.First(&photo, "id = ?", uuid.FromStringOrNil(photoId))
	return r.db.Model(&photo).Association("tags").Delete(tag).Error
}

func (r *PhotoRepository) ForUser(userId string, currentPage int64, itemsPerPage int64) *pagination.Paginator {
	var photos []models.Photo
	db := r.db.Preload("Tags").Preload("Apps").Where("user_id = ?", userId)
	return pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    int(currentPage),
		Limit:   int(itemsPerPage),
		OrderBy: []string{"created_at desc"},
		ShowSQL: r.debug,
	}, &photos)
}

func (r *PhotoRepository) ForTags(tags []string, currentPage int64, itemsPerPage int64) *pagination.Paginator {
	var photos []models.Photo
	db := r.db.
		Preload("Tags").
		Select("DISTINCT photos.*").
		Joins("left join photo_tag on photo_tag.photo_id=photos.id").
		Joins("left join tags on tags.id=photo_tag.tag_id").
		Where("tags.name IN (?)", tags)
	return pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    int(currentPage),
		Limit:   int(itemsPerPage),
		OrderBy: []string{"photos.created_at desc"},
		ShowSQL: r.debug,
	}, &photos)
}
