package action_handlers_test

import (
	action2 "git.jasonraimondi.com/jason/jasontest/domain/action"
	"git.jasonraimondi.com/jason/jasontest/domain/action_handlers"
	lib2 "git.jasonraimondi.com/jason/jasontest/domain/lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePersonHandler(t *testing.T) {
	r := lib2.NewTestApplication().RepositoryFactory()

	cp := action_handlers.NewCreatePersonHandler(r.Person())
	first := "Jason"
	last := "Raimondi"
	password := "jasonraimondi"
	err := cp.Handle(action2.NewCreatePerson(
		&first,
		&last,
		"jason@raimondi.us",
		&password,
	))
	assert.NoError(t, err)

	p, err := r.Person().GetByEmail("jason@raimondi.us")
	assert.NoError(t, err)
	assert.Equal(t, "jason@raimondi.us", p.Email)
	assert.Equal(t, "Jason", p.FirstName.String)
	assert.Equal(t, "Raimondi", p.LastName.String)
	assert.True(t, p.CheckPassword(password))
}
