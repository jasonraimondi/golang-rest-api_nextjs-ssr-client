package model_test

import (
	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerson_SetPassword(t *testing.T) {
	p := model.NewUser("jason@raimondi.us")
	password := "jasonraimondi"

	assert.NoError(t, p.SetPassword(password))
	assert.True(t, p.CheckPassword(password))
}
