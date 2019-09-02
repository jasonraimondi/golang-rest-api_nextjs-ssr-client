package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/server/responses"
)

type AppHandler struct {
	appRepository *repository.AppRepository
}

func (h *AppHandler) List(c echo.Context) error {
	page := strToInt(c.QueryParam("page"), 1)
	itemsPerPage := strToInt(c.QueryParam("itemsPerPage"), 25)
	paginator := h.appRepository.List(page, itemsPerPage)
	return responses.SendPaginator(c, http.StatusOK, paginator)
}
