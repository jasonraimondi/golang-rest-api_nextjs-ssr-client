package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type PhotoRepository struct {
	dbx *gorm.DB
}

func (r *PhotoRepository) GetById(id string) (photo models.Photo, err error) {
	photo = models.Photo{}
	err = r.dbx.First(photo).Error
	return photo, err
}

func (r *PhotoRepository) Update(u *models.Photo) (err error) {
	return r.dbx.Update(u).Error
}

func (r *PhotoRepository) Create(u *models.Photo) (err error) {
	return r.dbx.Create(u).Error
}

func (s *PhotoRepository) ForUser(userId string, currentPage int64, itemsPerPage int64) *pagination.Paginator {
	var photos []models.Photo
	db := s.dbx.Where("user_id = ?", userId)
	return pagination.Paging(&pagination.Param{
		DB: db,
		Page: int(currentPage),
		Limit: int(itemsPerPage),
		OrderBy: []string{"created_at desc"},
		ShowSQL: true,
	}, &photos)
}

func (s *PhotoRepository) ForTags(tags []string, currentPage int64, itemsPerPage int64) *pagination.Paginator {
	var photos []models.Photo
	db := s.dbx.Preload("tags").Joins("left join photo_tag on photo_tag.photo_id=photos.id").Joins("left join tags on tags.id=photo_tag.tag_id").Where("tags.name IN (?)", tags)
	return pagination.Paging(&pagination.Param{
		DB: db,
		Page: int(currentPage),
		Limit: int(itemsPerPage),
		OrderBy: []string{"photos.created_at desc"},
		ShowSQL: true,
	}, &photos)
}
