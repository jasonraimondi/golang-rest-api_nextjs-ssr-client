package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"git.jasonraimondi.com/jason/jasontest/models"
)

func TestUser_SetPassword(t *testing.T) {
	p := models.NewUser("jason@raimondi.us")
	password := "jasonraimondi"

	if assert.NoError(t, p.SetPassword(password)) {
		assert.True(t, p.CheckPassword(password))
	}
}

func TestUser_GetFullName(t *testing.T) {
	p := models.NewUser("jason1@raimondi.us")
	p.SetFirst("Jason")
	p.SetLast("Raimondi")

	p2 := models.NewUser("jason2@raimondi.us")
	p2.SetFirst("Jason")

	p3 := models.NewUser("jason3@raimondi.us")
	p3.SetLast("Raimondi")

	p4 := models.NewUser("jason4@raimondi.us")

	assert.Equal(t, "Jason Raimondi", p.GetFullName())
	assert.Equal(t, "Jason", p2.GetFullName())
	assert.Equal(t, "Raimondi", p3.GetFullName())
	assert.Equal(t, "", p4.GetFullName())
}

func TestUser_GetFullIdentifier(t *testing.T) {
	p := models.NewUser("jason@raimondi.us")
	p.SetFirst("Jason")
	p.SetLast("Raimondi")
	assert.Equal(t, "Jason Raimondi <jason@raimondi.us>", p.GetFullIdentifier())
}
