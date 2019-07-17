package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"gopkg.in/go-playground/validator.v9"

	"git.jasonraimondi.com/jason/jasontest/domain/repository"
)

type Service struct {
	repository *repository.Factory
	validate   *validator.Validate
	s3Config   *S3Config
}

func NewService(r *repository.Factory, v *validator.Validate, c *S3Config) *Service {
	return &Service{r, v, c}
}

type S3Config struct {
	OriginalBucket string
	*aws.Config
}

func NewS3Config(originalBucket string, aws *aws.Config) *S3Config {
	return &S3Config{
		OriginalBucket: originalBucket,
		Config: aws,
	}
}
