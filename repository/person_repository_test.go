package repository_test

import (
	"git.jasonraimondi.com/jason/jasontest/models"
	"git.jasonraimondi.com/jason/jasontest/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository(t *testing.T) {
	r, err := repository.Initialize()
	assert.NoError(t, err)
	p := models.NewPerson("jason@raimondi.us")
	err = r.Person().Create(p)
	assert.Equal(t, "jason@raimondi.us", p.Email)
}
