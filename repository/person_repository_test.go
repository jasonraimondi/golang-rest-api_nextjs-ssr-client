package repository_test

import (
	"git.jasonraimondi.com/jason/jasontest/models"
	"git.jasonraimondi.com/jason/jasontest/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPersonRepository_GetById(t *testing.T) {
	r := repository.NewTestApplication().RepositoryFactory()
	p := models.NewPerson("jason@raimondi.us")
	p.FirstName = models.ToNullString("Jason")
	p.LastName = models.ToNullString("Raimondi")
	err := r.Person().Create(p)
	assert.NoError(t, err)

	sut1, err := r.Person().GetById(p.ID)

	assert.NoError(t, err)
	assert.Equal(t, "jason@raimondi.us", sut1.Email)
	assert.Equal(t, "Jason", sut1.FirstName.String)
	assert.Equal(t, "Raimondi", sut1.LastName.String)
}

func TestPersonRepository_GetByEmail(t *testing.T) {
	r := repository.NewTestApplication().RepositoryFactory()
	p := models.NewPerson("kimberly@foo.bar")
	p.FirstName = models.ToNullString("Kimberly")
	p.LastName = models.ToNullString("Foo")
	err := r.Person().Create(p)
	assert.NoError(t, err)

	sut, err := r.Person().GetByEmail("kimberly@foo.bar")

	assert.NoError(t, err)
	assert.Equal(t, "kimberly@foo.bar", sut.Email)
	assert.Equal(t, "Kimberly", sut.FirstName.String)
	assert.Equal(t, "Foo", sut.LastName.String)
}