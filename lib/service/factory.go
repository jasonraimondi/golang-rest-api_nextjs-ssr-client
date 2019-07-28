package service

import (
	"gopkg.in/go-playground/validator.v9"

	"git.jasonraimondi.com/jason/jasontest/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/lib/s3"
)

type Factory struct {
	repository *repository.Factory
	validate   *validator.Validate
	s3         *s3.S3Config
}

func NewService(r *repository.Factory, v *validator.Validate, c *s3.S3Config) *Factory {
	return &Factory{r, v, c}
}

func (s *Factory) SignUpService() *SignUpService {
	return &SignUpService{
		repository:     s.repository,
		validate:       s.validate,
		userRepository: s.repository.User(),
	}
}

func (s *Factory) FileUploadService() *FileUploadService {
	return &FileUploadService{
		originals:      "originals",
		repository:     s.repository,
		userRepository: s.repository.User(),
		s3:             s.s3,
	}
}

func (s *Factory) ListPhotosService() *ListPhotosService {
	return &ListPhotosService{
		photoRepository: s.repository.PhotoRepository(),
	}
}
