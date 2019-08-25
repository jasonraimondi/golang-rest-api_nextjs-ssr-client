package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
)

type PhotoHandler struct {
	photoRepository    *repository.PhotoRepository
	photoUploadService *service.PhotoUploadService
	photoAppService    *service.PhotoAppService
}

func (h *PhotoHandler) ListForUser(c echo.Context) error {
	userId := c.Param("userId")

	page := strToInt(c.QueryParam("page"), 1)
	itemsPerPage := strToInt(c.QueryParam("itemsPerPage"), 25)

	paginator := h.photoRepository.ForUser(userId, page, itemsPerPage)
	return c.JSON(http.StatusOK, paginator)
}

func (h *PhotoHandler) ListForTags(c echo.Context) error {
	tags := c.QueryParams()["tags[]"]
	if tags == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "missing tags[]")
	}
	page := strToInt(c.QueryParam("page"), 1)
	itemsPerPage := strToInt(c.QueryParam("itemsPerPage"), 25)

	paginator := h.photoRepository.ForTags(tags, page, itemsPerPage)
	return c.JSON(http.StatusOK, paginator)
}

func (h *PhotoHandler) AttachTags(c echo.Context) error {
	photoId := c.Param("photoId")
	tags, _ := c.FormParams()
	if err := h.photoAppService.AddTagsToPhoto(photoId, tags["tags[]"]); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return sendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}

func (h *PhotoHandler) RemoveTag(c echo.Context) error {
	photoId := c.Param("photoId")
	if tagId, err := strconv.Atoi(c.Param("tagId")); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid tagId")
	} else if err = h.photoAppService.RemoveTagFromPhoto(photoId, uint(tagId)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return sendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}

func (h *PhotoHandler) Create(c echo.Context) error {
	userId := c.Param("userId")

	if form, err := c.MultipartForm(); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, "form error")
	} else if httpErr := h.photoUploadService.FileUpload(form, userId); httpErr != nil {
		return httpErr
	}

	return sendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}

func (h *PhotoHandler) Show(c echo.Context) error {
	userId := c.Param("photoId")
	photo, err := h.photoRepository.GetById(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, photo)
}
