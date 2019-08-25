package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
)

type TagHandler struct {
	tagRepository *repository.TagRepository
}

func (h *TagHandler) ListForPhoto(c echo.Context) error {
	photoId := c.Param("photoId")

	page := strToInt(c.QueryParam("page"), 1)
	itemsPerPage := strToInt(c.QueryParam("itemsPerPage"), 25)

	paginator := h.tagRepository.ForPhoto(photoId, page, itemsPerPage)
	return c.JSON(http.StatusOK, paginator)
}
