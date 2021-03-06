package handlers

import (
	"strconv"

	"git.jasonraimondi.com/jason/jasontest/app/lib"
)

type Handler struct {
	App *lib.Application
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
		tagService:         h.App.ServiceFactory.TagService(),
		photoService:       h.App.ServiceFactory.PhotoService(),
		photoUploadService: h.App.ServiceFactory.PhotoUploadService(),
	}
}

func (h *Handler) AppHandler() *AppHandler {
	return &AppHandler{
		appRepository: h.App.RepositoryFactory.AppRepository(),
	}
}

func (h *Handler) PhotoHandler() *PhotoHandler {
	return &PhotoHandler{
		photoRepository: h.App.RepositoryFactory.PhotoRepository(),
	}
}

func strToInt(s string, d int) int64 {
	i, err := strconv.Atoi(s)
	if err != nil {
		i = d
	}
	return int64(i)
}
