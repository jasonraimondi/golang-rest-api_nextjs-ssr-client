package repository

import (
	"github.com/jinzhu/gorm"
	paginator "github.com/pilagod/gorm-cursor-paginator"

	"git.jasonraimondi.com/jason/jasontest/app/lib/pagination"
	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type PhotoRepository struct {
	dbx *gorm.DB
}

func (r *PhotoRepository) GetById(id string) (photo Photo, err error) {
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

func (r *PhotoRepository) ListForUser(db *gorm.DB, q pagination.PagingQuery) ([]models.Photo, paginator.Cursor, error) {
	var photos []models.Photo

	stmt := db
	//stmt = db.Where(/* ... other filters ... */)
	//stmt = db.Or(/* ... more other filters ... */)

	// get paginator for photos.Photo
	p := pagination.GetPhotoPaginator(q)

	// use GORM-like syntax to do pagination
	result := p.Paginate(stmt, &photos)

	if result.Error != nil {
		//return nil, nil, result.Error
	}
	// get cursor for next iteration
	cursor := p.GetNextCursor()

	return photos, cursor, nil
}