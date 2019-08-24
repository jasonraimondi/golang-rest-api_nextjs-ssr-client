package repository

import (
	"errors"
	"fmt"
	"math"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/gorm"
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
	return query.
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

func TotalCountForQuery(dbx *gorm.DB, queryBuilder squirrel.SelectBuilder) (int64, error) {
	if sql, args, err := queryBuilder.Column("COUNT(*) as count").ToSql(); err != nil {
		return -1, err
	} else if rows, err := dbx.Raw(sql, args...).Rows(); err != nil {
		return -1, err
	} else {
		for rows.Next() {
			var count int64
			if err := dbx.ScanRows(rows, &count); err != nil {
				return -1, err
			}
			return count, nil
		}
	}
	return -1, errors.New("something went wrong counting the query")
}
