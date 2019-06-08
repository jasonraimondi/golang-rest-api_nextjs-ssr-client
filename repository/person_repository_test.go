package repository_test

import (
	"git.jasonraimondi.com/jason/jasontest/models"
	"git.jasonraimondi.com/jason/jasontest/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPersonRepository_GetById(t *testing.T) {
	r, err := repository.NewTestDB()
	assert.NoError(t, err)
	p1 := models.NewPerson("jason@raimondi.us")
	p1.FirstName = models.ToNullString("Jason")
	p1.LastName = models.ToNullString("Raimondi")

	if err = r.Person().Create(p1); err != nil {
		t.Fatalf("Error Creating (%s)", p1.Email)
	}

	sut1, err := r.Person().GetById(p1.ID)
	assert.NoError(t, err)
	assert.Equal(t, "jason@raimondi.us", sut1.Email)
	assert.Equal(t, "Jason", sut1.FirstName.String)
	assert.Equal(t, "Raimondi", sut1.LastName.String)
}

func TestPersonRepository_GetByEmail(t *testing.T) {
	r, err := repository.NewTestDB()
	assert.NoError(t, err)
	p := models.NewPerson("kimberly@foo.bar")
	p.FirstName = models.ToNullString("Kimberly")
	p.LastName = models.ToNullString("Foo")

	if err = r.Person().Create(p); err != nil {
		t.Fatalf("Error Creating (%s)", p.Email)
	}

	sut, err := r.Person().GetByEmail("kimberly@foo.bar")
	assert.NoError(t, err)
	assert.Equal(t, "kimberly@foo.bar", sut.Email)
	assert.Equal(t, "Kimberly", sut.FirstName.String)
	assert.Equal(t, "Foo", sut.LastName.String)
}