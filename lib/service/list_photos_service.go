package service

import (
	"fmt"
	"git.jasonraimondi.com/jason/jasontest/models"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"math"
)

type ListPhotosService struct {
	queryBuilder squirrel.StatementBuilderType
	dbx *sqlx.DB
}

func (s *ListPhotosService) ListPhotos(userId string, page int64, itemsPerPage int64) (*Paginator, error) {
	query := s.queryBuilder.Select().From("photos").Where(squirrel.Eq{
		"user_id": userId,
	})

	//photos, err := s.photoRepository.ListForUser(userId, int64(page), int64(itemsPerPage))
	//if err != nil {
	//	return nil, err
	//}
	//totalCount, err := s.photoRepository.CountForUser(userId)
	//if err != nil {
	//	return nil, err
	//}
	return s.createPaginator(query, itemsPerPage, page), nil
}

func (s *ListPhotosService) convertTypeToInterface(photos []models.Photo) []interface{} {
	results := make([]interface{}, len(photos))
	for i, v := range photos {
		results[i] = v
	}
	return results
}

type Paginator struct {
	Page         int64
	ItemsPerPage int64
	TotalCount   int64
	TotalPages   int64
	Data         []interface{}
	Links        Link
}

type Link struct {
	HasNextLink     bool
	HasPreviousLink bool
	Next            string
	Previous        string
}

func (s *ListPhotosService) createPaginator(queryBuilder squirrel.SelectBuilder, itemsPerPage int64, currentPage int64) *Paginator {
	totalCount, _ := s.GetTotalCount(queryBuilder)
	limit := itemsPerPage
	offset := limit * (currentPage - 1)

	queryBuilder = queryBuilder.Column("*").OrderBy("created_at DESC").Limit(uint64(limit)).Offset(uint64(offset))

	sql, args, err := queryBuilder.ToSql()

	if err != nil {
		return nil
	}

	data := []models.Photo{}
	err = s.dbx.Select(&data, sql, args...)
	if err != nil {
		return nil
	}



	totalPages := int64(math.Ceil(float64(totalCount) / float64(itemsPerPage)))
	var nextLink string
	var previousLink string
	hasNextLink := totalPages > currentPage
	hasPreviousLink := currentPage > 1
	if hasNextLink {
		nextLink = fmt.Sprintf("http://localhost:1323/list_photos?userId=3a38a226-3dd5-4d59-9694-6573c1e37cc1&itemsPerPage=%d&page=%d", itemsPerPage, currentPage + 1)
	}
	if hasPreviousLink {
		previousLink = fmt.Sprintf("http://localhost:1323/list_photos?userId=3a38a226-3dd5-4d59-9694-6573c1e37cc1&itemsPerPage=%d&page=%d", itemsPerPage, currentPage - 1)
	}
	links := Link{
		Next:            nextLink,
		Previous:        previousLink,
		HasNextLink:     hasNextLink,
		HasPreviousLink: hasPreviousLink,
	}
	paginator := &Paginator{
		Page:         currentPage,
		ItemsPerPage: itemsPerPage,
		Data:         s.convertTypeToInterface(data),
		TotalCount:   totalCount,
		TotalPages:   totalPages,
		Links:        links,
	}
	return paginator
}


func (r *ListPhotosService) GetTotalCount(queryBuilder squirrel.SelectBuilder) (count int64, err error) {
	sql, args, err := queryBuilder.Column("COUNT(*) as count").ToSql()
	if err != nil {
		return -1, err
	}
	rows, err := r.dbx.Query(sql, args...)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}
	return count, nil
}