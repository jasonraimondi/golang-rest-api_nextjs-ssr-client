package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (h *Handler) ListPhotos(c echo.Context) error {
	r := h.App.RepositoryFactory

	userId := c.QueryParam("userId")

	pp := c.QueryParam("page")
	page, err := strconv.Atoi(pp)
	if err != nil {
		page = 1
	}
	ipp := c.QueryParam("itemsPerPage")
	itemsPerPage, err := strconv.Atoi(ipp)
	if err != nil {
		itemsPerPage = 25
	}
	photos, err := r.PhotoRepository().ListForUser(userId, int64(page), int64(itemsPerPage))
	if err != nil {
		return err
	}
	s := make([]interface{}, len(photos))
	for i, v := range photos {
		s[i] = v
	}
	totalItems, err := r.PhotoRepository().CountForUser(userId)
	if err != nil {
		return err
	}
	return jsonItems(c, int64(page), int64(itemsPerPage), totalItems, s)
}

type PaginationItems struct {
	Page         int64
	ItemsPerPage int64
	TotalCount   int64
	Data         []interface{}
}

func jsonItems(c echo.Context, page int64, itemsPerPage int64, totalItems int64, data []interface{}) error {
	return c.JSON(http.StatusOK, &PaginationItems{
		Page:         page,
		ItemsPerPage: itemsPerPage,
		Data:         data,
		TotalCount:   totalItems,
	})
}
