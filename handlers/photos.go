package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (h *Handler) ListPhotos(c echo.Context) error {
	s := h.App.ServiceFactory

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
	paginator, err := s.ListPhotosService().ListPhotos(userId, int64(page), int64(itemsPerPage))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, paginator)
}
