package service

import (
	"gopkg.in/go-playground/validator.v9"

	"git.jasonraimondi.com/jason/jasontest/app/lib/config"
	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
)

type Factory struct {
	repository   *repository.Factory
	validate     *validator.Validate
	s3           *config.S3Config
	jwtSecureKey config.JWTSecureKey
}

func NewService(r *repository.Factory, v *validator.Validate, c *config.S3Config, j config.JWTSecureKey) *Factory {
	return &Factory{r, v, c, j}
}

func (s *Factory) SignUpService() *SignUpService {
	return &SignUpService{
		validate:                     s.validate,
		signUpConfirmationRepository: s.repository.SignUpConfirmation(),
		userRepository:               s.repository.UserRepository(),
	}
}

func (s *Factory) FileUploadService() *PhotoUploadService {
	return &PhotoUploadService{
		originals:      BucketName("originals"),
		repository:     s.repository,
		userRepository: s.repository.UserRepository(),
		s3:             s.s3,
	}
}

func (s *Factory) AuthService() *AuthService {
	return &AuthService{
		userRepository: s.repository.UserRepository(),
		jwtSecureKey:   s.jwtSecureKey,
	}
}

func (s *Factory) PhotoAppService() *TagService {
	return &TagService{
		db:              s.repository.DB(),
		photoRepository: s.repository.PhotoRepository(),
	}
}

//func (s *Factory) PhotoTagService() *PhotoTagService {
//	return &PhotoTagService{
//		repository: s.repository,
//	}
//}
