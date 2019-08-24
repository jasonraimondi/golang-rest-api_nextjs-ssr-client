package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
)

type PhotoHandler struct {
	//listTagsRepository   *repository.ListTagsRepository
	//listAppsRepository   *repository.ListAppsRepository
	//listPhotosRepository *repository.ListPhotosRepository
	photoUploadService   *service.PhotoUploadService
	photoAppService      *service.PhotoAppService
}

//func (h *PhotoHandler) ListForUser(c echo.Context) error {
//	userId := c.Param("userId")
//
//	page := strToInt(c.QueryParam("page"), 1)
//	itemsPerPage := strToInt(c.QueryParam("itemsPerPage"), 25)
//
//	paginator, err := h.listPhotosRepository.ForUser(userId, page, itemsPerPage)
//	if err != nil {
//		return err
//	}
//	return c.JSON(http.StatusOK, paginator)
//}
//
//func (h *PhotoHandler) ListForTags(c echo.Context) error {
//	tagNames := c.QueryParams()["tagNames[]"]
//	if tagNames == nil {
//		return echo.NewHTTPError(http.StatusBadRequest, "missing tagNames[]")
//	}
//	page := strToInt(c.QueryParam("page"), 1)
//	itemsPerPage := strToInt(c.QueryParam("itemsPerPage"), 25)
//
//	paginator, err := h.listPhotosRepository.ForTags(tagNames, page, itemsPerPage)
//	if err != nil {
//		return err
//	}
//	return c.JSON(http.StatusOK, paginator)
//}
//
//// Move To Apps Handler
//func (h *PhotoHandler) ListApps(c echo.Context) error {
//	photoId := c.Param("photoId")
//
//	page := strToInt(c.QueryParam("page"), 1)
//	itemsPerPage := strToInt(c.QueryParam("itemsPerPage"), 25)
//
//	paginator, err := h.listAppsRepository.ForPhoto(photoId, page, itemsPerPage)
//	if err != nil {
//		return err
//	}
//	return c.JSON(http.StatusOK, paginator)
//}
//
//func (h *PhotoHandler) ListTags(c echo.Context) error {
//	photoId := c.Param("photoId")
//
//	page := strToInt(c.QueryParam("page"), 1)
//	itemsPerPage := strToInt(c.QueryParam("itemsPerPage"), 25)
//
//	paginator, err := h.listTagsRepository.ForPhoto(photoId, page, itemsPerPage)
//	if err != nil {
//		return err
//	}
//	return c.JSON(http.StatusOK, paginator)
//}

func (h *PhotoHandler) LinkTags(c echo.Context) error {
	photoId := c.Param("photoId")
	tags, _ := c.FormParams()
	if err := h.photoAppService.AddTagsToPhoto(photoId, tags["tags[]"]); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return sendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}

func (h *PhotoHandler) RemoveApp(c echo.Context) error {
	photoId := c.Param("photoId")
	appId, err := strconv.ParseInt(c.Param("appId"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid appId")
	}
	if err = h.photoAppService.RemoveAppFromPhoto(photoId, appId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return sendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}

func (h *PhotoHandler) RemoveTag(c echo.Context) error {
	photoId := c.Param("photoId")
	tagId, err := strconv.ParseInt(c.Param("tagId"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid tagId")
	}
	if err := h.photoAppService.RemoveTagFromPhoto(photoId, tagId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return sendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}

func (h *PhotoHandler) LinkApps(c echo.Context) error {
	photoId := c.Param("photoId")
	apps, _ := c.FormParams()
	if err := h.photoAppService.AddAppsToPhoto(photoId, apps["apps[]"]); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return sendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}

func (h *PhotoHandler) Create(c echo.Context) error {
	// Read form fields
	userId := c.FormValue("userId")

	if form, err := c.MultipartForm(); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, "form error")
	} else if httpErr := h.photoUploadService.FileUpload(form, userId); httpErr != nil {
		return httpErr
	}

	return sendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}
