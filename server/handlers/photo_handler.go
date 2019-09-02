package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/server/responses"
)

type PhotoHandler struct {
	photoRepository *repository.PhotoRepository
}

func (h *PhotoHandler) ListForUser(c echo.Context) error {
	userId := c.Param("userId")

	page := strToInt(c.QueryParam("page"), 1)
	itemsPerPage := strToInt(c.QueryParam("itemsPerPage"), 25)

	paginator := h.photoRepository.ForUser(userId, page, itemsPerPage)
	return responses.SendPaginator(c, http.StatusOK, paginator)
}

func (h *PhotoHandler) ListForApp(c echo.Context) error {
	appId := c.Param("appId")

	page := strToInt(c.QueryParam("page"), 1)
	itemsPerPage := strToInt(c.QueryParam("itemsPerPage"), 25)

	paginator := h.photoRepository.ForApp(appId, page, itemsPerPage)
	return responses.SendPaginator(c, http.StatusOK, paginator)
}

func (h *PhotoHandler) ListForTags(c echo.Context) error {
	tags := c.QueryParams()["tags[]"]
	if tags == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "missing tags[]")
	}
	page := strToInt(c.QueryParam("page"), 1)
	itemsPerPage := strToInt(c.QueryParam("itemsPerPage"), 25)

	paginator := h.photoRepository.ForTags(tags, page, itemsPerPage)
	return responses.SendPaginator(c, http.StatusOK, paginator)
}

func (h *PhotoHandler) Show(c echo.Context) error {
	photoId := c.Param("photoId")
	photo, err := h.photoRepository.GetById(photoId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return responses.SendAny(c, http.StatusOK, photo)
}
