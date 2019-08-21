package repository

import (
	"fmt"
	"math"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Paginator struct {
	Page         int64
	ItemsPerPage int64
	TotalCount   int64
	TotalPages   int64
	Data         []interface{}
	Link         Link
}

type Link struct {
	HasNext     bool
	HasPrevious bool
	Next        string
	Previous    string
}

func PaginateQuery(itemsPerPage int64, currentPage int64, query squirrel.SelectBuilder) (string, []interface{}, error) {
	limit := itemsPerPage
	offset := limit * (currentPage - 1)

	return query.Column("*").
		OrderBy("created_at DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
}

func NewPaginator(baseURL string, totalCount int64, itemsPerPage int64, currentPage int64, data []interface{}) (*Paginator, error) {
	var nextLink string
	var previousLink string

	totalPages := int64(math.Ceil(float64(totalCount) / float64(itemsPerPage)))
	hasNextLink := totalPages > currentPage
	hasPreviousLink := currentPage > 1

	if hasNextLink {
		nextLink = fmt.Sprintf("%s&itemsPerPage=%d&page=%d", baseURL, itemsPerPage, currentPage+1)
	}
	if hasPreviousLink {
		previousLink = fmt.Sprintf("%s&itemsPerPage=%d&page=%d", baseURL, itemsPerPage, currentPage-1)
	}
	if len(data) == 0 {
		data = []interface{}{}
	}

	links := Link{
		Next:        nextLink,
		Previous:    previousLink,
		HasNext:     hasNextLink,
		HasPrevious: hasPreviousLink,
	}
	paginator := &Paginator{
		Page:         currentPage,
		ItemsPerPage: itemsPerPage,
		Data:         data,
		TotalCount:   totalCount,
		TotalPages:   totalPages,
		Link:         links,
	}
	return paginator, nil
}

func TotalCountForQuery(dbx *sqlx.DB, queryBuilder squirrel.SelectBuilder) (count int64, err error) {
	if sql, args, err := queryBuilder.Column("COUNT(*) as count").ToSql(); err != nil {
		return -1, err
	} else if rows, err := dbx.Query(sql, args...); err != nil {
		return -1, err
	} else {
		for rows.Next() {
			if err := rows.Scan(&count); err != nil {
				return -1, err
			}
		}
	}
	return count, nil
}
