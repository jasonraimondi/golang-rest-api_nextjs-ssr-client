package repository_test

import (
	"git.jasonraimondi.com/jason/jasontest/models"
	"git.jasonraimondi.com/jason/jasontest/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_GetById(t *testing.T) {
	r, err := repository.Initialize()
	assert.NoError(t, err)
	p := models.NewPerson("jason@raimondi.us")
	p.FirstName = models.ToNullString("Jason")
	p.LastName = models.ToNullString("Raimondi")
	err = r.Person().Create(p)
	assert.NoError(t, err)

	sut, err := r.Person().GetById(p.Id)
	assert.NoError(t, err)
	assert.Equal(t, "jason@raimondi.us", sut.Email)
	assert.Equal(t, "Jason", sut.FirstName.String)
	assert.Equal(t, "Raimondi", sut.LastName.String)
}
