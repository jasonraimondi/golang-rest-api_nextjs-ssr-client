package service

import (
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
	"gopkg.in/go-playground/validator.v9"
)

type Service struct {
	repository *repository.Factory
	validate   *validator.Validate
}

func NewService(r *repository.Factory, v *validator.Validate) *Service {
	return &Service{r, v}
}
