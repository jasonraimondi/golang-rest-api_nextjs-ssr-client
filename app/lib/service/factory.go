package service

import (
	"gopkg.in/go-playground/validator.v9"

	"git.jasonraimondi.com/jason/jasontest/app/lib/awsupload"
	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
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

func (s *Factory) FileUploadService() *PhotoUploadService {
	return &PhotoUploadService{
		originals:      "originals",
		repository:     s.repository,
		userRepository: s.repository.User(),
		s3:             s.s3,
	}
}

func (s *Factory) AuthService() *AuthService {
	return &AuthService{
		userRepository: s.repository.User(),
		jwtSecureKey:   s.jwtSecureKey,
	}
}

func (s *Factory) TagService() *TagService {
	return &TagService{
		repository: s.repository,
	}
}

