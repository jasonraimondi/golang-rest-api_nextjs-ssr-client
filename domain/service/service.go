package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"gopkg.in/go-playground/validator.v9"

	"git.jasonraimondi.com/jason/jasontest/domain/repository"
)

type Service struct {
	repository *repository.Factory
	validate   *validator.Validate
	s3Config   *aws.Config
}

func NewService(r *repository.Factory, v *validator.Validate, c *aws.Config) *Service {
	return &Service{r, v, c}
}
