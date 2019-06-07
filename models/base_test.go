package models_test

import (
	"git.jasonraimondi.com/jason/jasontest/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNullString(t *testing.T) {
	isNullString := models.ToNullString("")
	isValidString := models.ToNullString("jason")
	assert.False(t, isNullString.Valid)
	assert.True(t, isValidString.Valid)
}
