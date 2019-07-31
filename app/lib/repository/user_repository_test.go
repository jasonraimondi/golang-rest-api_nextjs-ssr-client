package repository_test

import (
	"git.jasonraimondi.com/jason/jasontest/app/models"
	"git.jasonraimondi.com/jason/jasontest/app/test/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersonRepository_GetById(t *testing.T) {
	r := utils.NewTestApplication().RepositoryFactory
	u := models.NewUser("jason@raimondi.us")
	u.SetFirst("Jason")
	u.SetLast("Raimondi")
	err := r.User().Create(u)
	assert.NoError(t, err)

	sut1, err := r.User().GetById(u.GetID().String())

	if assert.NoError(t, err) {
		assert.Equal(t, "jason@raimondi.us", sut1.Email)
		assert.Equal(t, "Jason", sut1.First.String)
		assert.Equal(t, "Raimondi", sut1.Last.String)
	}
}

func TestPersonRepository_GetByEmail(t *testing.T) {
	r := utils.NewTestApplication().RepositoryFactory
	u := models.NewUser("kimberly@foo.bar")
	u.SetFirst("Kimberly")
	u.SetLast("Foo")
	err := r.User().Create(u)
	assert.NoError(t, err)

	sut, err := r.User().GetByEmail("kimberly@foo.bar")

	if assert.NoError(t, err) {
		assert.Equal(t, "kimberly@foo.bar", sut.Email)
		assert.Equal(t, "Kimberly", sut.First.String)
		assert.Equal(t, "Foo", sut.Last.String)
	}
}
