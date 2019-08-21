package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type ListPhotosRepository struct {
	queryBuilder squirrel.StatementBuilderType
	dbx          *sqlx.DB
}

func (s *ListPhotosRepository) ForUser(userId string, currentPage int64, itemsPerPage int64) (*Paginator, error) {
	query := s.queryBuilder.
		Select().
		From("photos").
		Where(squirrel.Eq{
			"user_id": userId,
		}).
		OrderBy("created_at DESC")

	totalCount, err := TotalCountForQuery(s.dbx, query)
	if err != nil {
		return nil, err
	}
	query = query.Column("*")
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
	// @todo env var for base url
	baseURL := "http://localhost:1323/photos/user/" + userId
	return NewPaginator(baseURL, totalCount, itemsPerPage, currentPage, results)
}
