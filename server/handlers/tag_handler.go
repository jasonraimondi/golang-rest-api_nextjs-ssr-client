package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
)

type TagHandler struct {
	tagService *service.TagService
}

func (h *TagHandler) Tag(c echo.Context) error {
	photoId := c.Param("photoId")
	//tags := []string{"one", "two", "foo", "bar"}

	tags2, _ := c.FormParams()

	tags := tags2["tag[]"]

	err := h.tagService.AddTagsToPhoto(photoId, tags)
	if err != nil {
	return err;
	}
	return sendMessage(c, http.StatusAccepted, http.StatusText(http.StatusAccepted))
}
