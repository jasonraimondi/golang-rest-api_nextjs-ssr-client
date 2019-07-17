package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"git.jasonraimondi.com/jason/jasontest/domain/lib"
	"git.jasonraimondi.com/jason/jasontest/domain/model"
)

func TestPersonRepository_GetById(t *testing.T) {
	r := lib.NewTestApplication().RepositoryFactory
	u := model.NewUser("jason@raimondi.us")
	u.First = model.ToNullString("Jason")
	u.Last = model.ToNullString("Raimondi")
	err := r.User().Create(u)
	assert.NoError(t, err)

	sut1, err := r.User().GetById(u.GetID())

	if assert.NoError(t, err) {
		assert.Equal(t, "jason@raimondi.us", sut1.Email)
		assert.Equal(t, "Jason", sut1.First.String)
		assert.Equal(t, "Raimondi", sut1.Last.String)
	}
}

func TestPersonRepository_GetByEmail(t *testing.T) {
	r := lib.NewTestApplication().RepositoryFactory
	u := model.NewUser("kimberly@foo.bar")
	u.First = model.ToNullString("Kimberly")
	u.Last = model.ToNullString("Foo")
	err := r.User().Create(u)
	assert.NoError(t, err)

	sut, err := r.User().GetByEmail("kimberly@foo.bar")

	if assert.NoError(t, err) {
		assert.Equal(t, "kimberly@foo.bar", sut.Email)
		assert.Equal(t, "Kimberly", sut.First.String)
		assert.Equal(t, "Foo", sut.Last.String)
	}
}
