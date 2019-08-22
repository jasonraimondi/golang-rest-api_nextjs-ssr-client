package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type ListTagsRepository struct {
	queryBuilder squirrel.StatementBuilderType
	dbx          *sqlx.DB
}

func (s *ListTagsRepository) ForPhoto(photoId string, currentPage int64, itemsPerPage int64) (*Paginator, error) {
	query := s.queryBuilder.Select().
		From("tags").
		Join("photo_tag ON photo_tag.tag_id=tags.id").
		Where(squirrel.Eq{
			"photo_tag.photo_id": photoId,
		})
	totalCount, err := TotalCountForQuery(s.dbx, query)
	query = query.Column("tags.*")
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
		p := models.Tag{}
		err = rows.StructScan(&p)
		if err != nil {
			continue
		}
		results = append(results, p)
	}
	// @todo env var for base url
	baseURL := "http://localhost:1323/photos/" + photoId + "/tags"
	return NewPaginator(baseURL, totalCount, itemsPerPage, currentPage, results)
}