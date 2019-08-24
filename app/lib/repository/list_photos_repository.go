package repository

import (
	"strings"

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
		})
	baseURL := "http://localhost:1323/photos/user/" + userId
	return s.listPhotos(query, "*", itemsPerPage, currentPage, baseURL)
}

func (s *ListPhotosRepository) ForTags(tagNames []string, currentPage int64, itemsPerPage int64) (*Paginator, error) {
	query := s.queryBuilder.
		Select().
		From("photos").
		LeftJoin("photo_tag on photo_tag.photo_id=photos.id").
		LeftJoin("tags on photo_tag.tag_id=tags.id").
		Where(squirrel.Eq{
			"tags.name": tagNames,
		})
	baseURL := "http://localhost:1323/photos/user/" + strings.Join(tagNames, " ")
	subq := s.queryBuilder.
		Select("tags.*").
		From("tags").
		LeftJoin("photo_tag on photo_tag.tag_id=tags.id").
		Where(squirrel.Eq{
			"photo_tag.photo_id":
		})
	return s.listPhotos(query, "photos.*", itemsPerPage, currentPage, baseURL)
}

func (s *ListPhotosRepository) listPhotos(query squirrel.SelectBuilder, sel string, itemsPerPage int64, currentPage int64, baseURL string) (*Paginator, error) {
	totalCount, err := TotalCountForQuery(s.dbx, query)
	if err != nil {
		return nil, err
	}
	query = query.Column(sel).OrderBy("created_at DESC")
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
	return NewPaginator(baseURL, totalCount, itemsPerPage, currentPage, results)
}
