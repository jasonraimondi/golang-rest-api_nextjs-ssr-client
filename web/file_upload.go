package web

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) Upload(c echo.Context) error {
	// Read form fields
	userId := c.FormValue("userId")

	if form, err := c.MultipartForm(); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, "form error")
	} else if httpErr := h.App.ServiceFactory.FileUpload(form, userId); httpErr != nil {
		return httpErr
	}

	return sendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}

func sendMessage(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, map[string]interface{}{
		"message": message,
	})
}
