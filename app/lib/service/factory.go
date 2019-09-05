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

func (s *Factory) PhotoUploadService() *PhotoUploadService {
	return &PhotoUploadService{
		bucketName:     s.s3.OriginalBucket,
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

func (s *Factory) PhotoService() *PhotoService {
	return &PhotoService{
		db:              s.repository.DB(),
		photoRepository: s.repository.PhotoRepository(),
		tagService:      s.TagService(),
	}
}

func (s *Factory) TagService() *TagService {
	return &TagService{
		db:              s.repository.DB(),
		photoRepository: s.repository.PhotoRepository(),
	}
}
