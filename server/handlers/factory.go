package handlers

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/labstack/echo"

	"git.jasonraimondi.com/jason/jasontest/app/lib"
)

type Handler struct {
	App      *lib.Application
	S3Config *aws.Config
}

func NewHandler(a *lib.Application) *Handler {
	return &Handler{
		App: a,
	}
}

func (h *Handler) SignUp() *SignUpHandler {
	return &SignUpHandler{
		signUp: h.App.ServiceFactory.SignUpService(),
	}
}

func (h *Handler) Auth() *AuthHandler {
	return &AuthHandler{
		factory: h.App.ServiceFactory,
	}
}

func (h *Handler) Tag() *TagHandler {
	return &TagHandler{
		tagService: h.App.ServiceFactory.TagService(),
	}
}

func (h *Handler) Photo() *PhotoHandler {
	return &PhotoHandler{
		listTagService:   h.App.RepositoryFactory.ListTagsRepository(),
		listPhotoService: h.App.RepositoryFactory.ListPhotosRepository(),
		photoUpload:      h.App.ServiceFactory.FileUploadService(),
	}
}

func sendMessage(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, map[string]interface{}{
		"message": message,
	})
}

func strToInt(s string, d int) int64 {
	i, err := strconv.Atoi(s)
	if err != nil {
		i = d
	}
	return int64(i)
}
