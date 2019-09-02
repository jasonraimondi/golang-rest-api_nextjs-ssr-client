package repository

import (
	"github.com/jinzhu/gorm"

	"git.jasonraimondi.com/jason/jasontest/app/lib/pagination"
	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type AppRepository struct {
	debug bool
	db    *gorm.DB
}

func (r *AppRepository) List(currentPage int64, itemsPerPage int64) *pagination.Paginator {
	var photos []models.App
	db := r.db.Where("id > ?", 0)
	return pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    int(currentPage),
		Limit:   int(itemsPerPage),
		OrderBy: []string{"name asc"},
		ShowSQL: r.debug,
	}, &photos)
}
