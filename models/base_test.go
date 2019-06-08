package models_test

import (
	"git.jasonraimondi.com/jason/jasontest/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNullString(t *testing.T) {
	invalid := models.ToNullString("")
	valid := models.ToNullString("jason")

	assert.False(t, invalid.Valid)
	assert.True(t, valid.Valid)
}
