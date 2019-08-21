package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
)

type PhotoHandler struct {
	listPhotoService *repository.ListPhotosRepository
	photoUpload      *service.PhotoUploadService
}

func (h *PhotoHandler) List(c echo.Context) error {

	userId := c.QueryParam("userId")

	page := strToInt(c.QueryParam("page"), 1)
	itemsPerPage := strToInt(c.QueryParam("itemsPerPage"), 25)

	paginator, err := h.listPhotoService.ForUser(userId, int64(page), int64(itemsPerPage))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, paginator)
}

func (h *PhotoHandler) Upload(c echo.Context) error {
	// Read form fields
	userId := c.FormValue("userId")

	if form, err := c.MultipartForm(); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, "form error")
	} else if httpErr := h.photoUpload.FileUpload(form, userId); httpErr != nil {
		return httpErr
	}

	return sendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}
