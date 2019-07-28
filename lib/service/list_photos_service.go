package service

import (
	"git.jasonraimondi.com/jason/jasontest/models"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type ListPhotosService struct {
	queryBuilder squirrel.StatementBuilderType
	dbx          *sqlx.DB
}

func (s *ListPhotosService) ListPhotos(userId string, currentPage int64, itemsPerPage int64) (*Paginator, error) {
	query := s.queryBuilder.Select().From("photos").Where(squirrel.Eq{
		"user_id": userId,
	})
	totalCount, err := TotalCountForQuery(s.dbx, query)
	if err != nil {
		return nil, err
	}
	sql, args, err := PaginateQuery(itemsPerPage, currentPage, query)
	if err != nil {
		return nil, err
	}
	rows, err := s.dbx.Queryx(sql, args...)
	if err != nil {
		return nil, err
	}
	var results []interface{}
	for rows.Next() {
		p := models.Photo{}
		err = rows.StructScan(&p)
		if err != nil {
			continue
		}
		results = append(results, p)
	}
	baseURL := "http://localhost:1323/list_photos?userId="+userId
	return NewPaginator(baseURL, totalCount, itemsPerPage, currentPage, results)
}
