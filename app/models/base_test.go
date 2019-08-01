package models_test

import (
	"testing"

	"git.jasonraimondi.com/jason/jasontest/app/models"

	"github.com/stretchr/testify/assert"
)

func TestToNullString(t *testing.T) {
	invalid := models.ToNullString("")
	valid := models.ToNullString("jason")

	assert.False(t, invalid.Valid)
	assert.True(t, valid.Valid)
}

func TestToNullInt64(t *testing.T) {
	invalid := models.ToNullInt64("")
	valid := models.ToNullInt64("634")

	assert.False(t, invalid.Valid)
	assert.True(t, valid.Valid)
}
