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
	tags := []string{"one", "two", "foo", "bar"}
	err := h.tagService.AddTagsToPhoto(tags)
	if err != nil {
		return err;
	}
	return sendMessage(c, http.StatusCreated, http.StatusText(http.StatusCreated))
}
