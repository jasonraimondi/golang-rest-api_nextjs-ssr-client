package action_test

import (
	"git.jasonraimondi.com/jason/jasontest/action"
	"git.jasonraimondi.com/jason/jasontest/lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePerson(t *testing.T) {
	r := lib.NewTestApplication().RepositoryFactory()

	cp := &action.CreatePersonHandler{
		PersonRepository: r.Person(),
	}
	first := "Jason"
	last := "Raimondi"
	password := "jasonraimondi"
	err := cp.Handle(&action.CreatePerson{
		First:    &first,
		Last:     &last,
		Email:    "jason@raimondi.us",
		Password: &password,
	})
	assert.NoError(t, err)

	p, err := r.Person().GetByEmail("jason@raimondi.us")
	assert.NoError(t, err)
	assert.Equal(t, "jason@raimondi.us", p.Email)
	assert.Equal(t, "Jason", p.FirstName.String)
	assert.Equal(t, "Raimondi", p.LastName.String)
	assert.True(t, p.CheckPassword(password))
}
