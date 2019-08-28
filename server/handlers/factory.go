package handlers

import (
	"strconv"

	"git.jasonraimondi.com/jason/jasontest/app/lib"
)

type Handler struct {
	App      *lib.Application
}

func NewHandler(a *lib.Application) *Handler {
	return &Handler{
		App: a,
	}
}

func (h *Handler) SignUpHandler() *SignUpHandler {
	return &SignUpHandler{
		signUp: h.App.ServiceFactory.SignUpService(),
	}
}

func (h *Handler) AuthHandler() *AuthHandler {
	return &AuthHandler{
		factory: h.App.ServiceFactory,
	}
}

func (h *Handler) AdminPhoto() *AdminPhotoHandler {
	return &AdminPhotoHandler{
		photoAppService:    h.App.ServiceFactory.PhotoAppService(),
		photoUploadService: h.App.ServiceFactory.FileUploadService(),
	}
}

func (h *Handler) Photo() *PhotoHandler {
	return &PhotoHandler{
		photoRepository:    h.App.RepositoryFactory.PhotoRepository(),
	}
}

func strToInt(s string, d int) int64 {
	i, err := strconv.Atoi(s)
	if err != nil {
		i = d
	}
	return int64(i)
}
