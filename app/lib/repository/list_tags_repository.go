package repository

import (
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type ListTagsRepository struct {
	queryBuilder squirrel.StatementBuilderType
	dbx          *sqlx.DB
}


func (s *ListTagsRepository) ForPhoto(photoId string, currentPage int64, itemsPerPage int64) (*Paginator, error) {
	//query := s.queryBuilder.Select().From("photos").Where(squirrel.Eq{
	//	"photo_id": photoId,
	//})
	// @todo wip add join to table
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
	// @todo env var for base url
	baseURL := "http://localhost:1323/list_tags?names=" + strings.Join(names, ",")
	return NewPaginator(baseURL, totalCount, itemsPerPage, currentPage, results)
}


func (s *ListTagsRepository) ForNames(names []string, currentPage int64, itemsPerPage int64) (*Paginator, error) {
	query := s.queryBuilder.Select().From("names").Where(squirrel.Eq{
		"name": names,
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
	// @todo env var for base url
	baseURL := "http://localhost:1323/list_tags?names=" + strings.Join(names, ",")
	return NewPaginator(baseURL, totalCount, itemsPerPage, currentPage, results)
}
