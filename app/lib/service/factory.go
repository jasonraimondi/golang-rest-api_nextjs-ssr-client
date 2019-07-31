package service

import (
	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/lib/awsupload"
	"github.com/Masterminds/squirrel"
	"gopkg.in/go-playground/validator.v9"
)

type Factory struct {
	repository   *repository.Factory
	validate     *validator.Validate
	s3           *awsupload.S3Config
	jwtSecureKey string
}

func NewService(r *repository.Factory, v *validator.Validate, c *awsupload.S3Config, j string) *Factory {
	return &Factory{r, v, c, j}
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
		queryBuilder: s.getPGQueryBuilder(),
		dbx:          s.repository.DBx,
	}
}

func (s *Factory) AuthService() *AuthService {
	return &AuthService{
		userRepository: s.repository.User(),
		jwtSecureKey:   s.jwtSecureKey,
	}
}

func (s *Factory) getPGQueryBuilder() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
}
