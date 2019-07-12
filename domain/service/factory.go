package service

import (
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
)

type Factory struct {
	r *repository.Factory
}

func NewFactory(r *repository.Factory) *Factory {
	return &Factory{r}
}
