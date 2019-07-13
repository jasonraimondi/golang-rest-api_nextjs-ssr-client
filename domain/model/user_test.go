package model_test

import (
	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_SetPassword(t *testing.T) {
	p := model.NewUser("jason@raimondi.us")
	password := "jasonraimondi"

	if assert.NoError(t, p.SetPassword(password)) {
		assert.True(t, p.CheckPassword(password))
	}
}

func TestUser_GetFullName(t *testing.T) {
	p := model.NewUser("jason1@raimondi.us")
	p.FirstName = model.ToNullString("Jason")
	p.LastName = model.ToNullString("Raimondi")

	p2 := model.NewUser("jason2@raimondi.us")
	p2.FirstName = model.ToNullString("Jason")

	p3 := model.NewUser("jason3@raimondi.us")
	p3.LastName = model.ToNullString("Raimondi")

	p4 := model.NewUser("jason4@raimondi.us")

	assert.Equal(t, "Jason Raimondi", p.GetFullName())
	assert.Equal(t, "Jason", p2.GetFullName())
	assert.Equal(t, "Raimondi", p3.GetFullName())
	assert.Equal(t, "", p4.GetFullName())
}

func TestUser_GetFullIdentifier(t *testing.T) {
	p := model.NewUser("jason@raimondi.us")
	p.FirstName = model.ToNullString("Jason")
	p.LastName = model.ToNullString("Raimondi")
	assert.Equal(t, "Jason Raimondi <jason@raimondi.us>", p.GetFullIdentifier())
}
