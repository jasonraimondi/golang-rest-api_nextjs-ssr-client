package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
	"git.jasonraimondi.com/jason/jasontest/server/responses"
)

type AdminPhotoHandler struct {
	photoUploadService *service.PhotoUploadService
	photoService       *service.PhotoService
	tagService         *service.TagService
}

func (h *AdminPhotoHandler) UpdatePhoto(c echo.Context) error {
	photoId := c.Param("photoId")
	p, _ := c.FormParams()
	tags := p["tags[]"]
	app := c.FormValue("app")
	description := c.FormValue("description")
	if err := h.photoService.UpdatePhoto(photoId, description, app, tags); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return responses.SendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}

//func (h *AdminPhotoHandler) AttachTags(c echo.Context) error {
//	photoId := c.Param("photoId")
//	tags, _ := c.FormParams()
//	if err := h.tagService.AddTagsToPhoto(photoId, tags["tags[]"]); err != nil {
//		return echo.NewHTTPError(http.StatusInternalServerError, err)
//	}
//	return responses.SendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
//}

//func (h *AdminPhotoHandler) RemoveApp(c echo.Context) error {
//	photoId := c.Param("photoId")
//	if appId, err := strconv.Atoi(c.Param("appId")); err != nil {
//		return echo.NewHTTPError(http.StatusBadRequest, "invalid appId")
//	} else if err = h.tagService.RemoveAppFromPhoto(photoId, uint(appId)); err != nil {
//		return echo.NewHTTPError(http.StatusInternalServerError, err)
//	}
//	return responses.SendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
//}

func (h *AdminPhotoHandler) RemoveTag(c echo.Context) error {
	photoId := c.Param("photoId")
	if tagId, err := strconv.Atoi(c.Param("tagId")); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid tagId")
	} else if err = h.tagService.RemoveTagFromPhoto(photoId, uint(tagId)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return responses.SendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}

func (h *AdminPhotoHandler) Create(c echo.Context) error {
	userId := c.Param("userId")

	if form, err := c.MultipartForm(); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, "form error")
	} else if httpErr := h.photoUploadService.FileUpload(form, userId); httpErr != nil {
		return httpErr
	}

	return responses.SendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}
