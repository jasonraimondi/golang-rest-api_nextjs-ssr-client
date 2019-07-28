package service

import (
	"fmt"
	"git.jasonraimondi.com/jason/jasontest/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/models"
	"math"
)

type ListPhotosService struct {
	photoRepository *repository.PhotoRepository
}

func (s *ListPhotosService) ListPhotos(userId string, page int64, itemsPerPage int64) (*Paginator, error) {
	photos, err := s.photoRepository.ListForUser(userId, int64(page), int64(itemsPerPage))
	if err != nil {
		return nil, err
	}
	totalCount, err := s.photoRepository.CountForUser(userId)
	if err != nil {
		return nil, err
	}
	return createPaginator(totalCount, itemsPerPage, page, s.convertTypeToInterface(photos)), nil
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

func createPaginator(totalCount int64, itemsPerPage int64, currentPage int64, data []interface{}) *Paginator {
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
		Data:         data,
		TotalCount:   totalCount,
		TotalPages:   totalPages,
		Links:        links,
	}
	return paginator
}
