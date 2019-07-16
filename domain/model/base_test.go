package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"git.jasonraimondi.com/jason/jasontest/domain/model"
)

func TestToNullString(t *testing.T) {
	invalid := model.ToNullString("")
	valid := model.ToNullString("jason")

	assert.False(t, invalid.Valid)
	assert.True(t, valid.Valid)
}

func TestToNullInt64(t *testing.T) {
	invalid := model.ToNullInt64("")
	valid := model.ToNullInt64("634")

	assert.False(t, invalid.Valid)
	assert.True(t, valid.Valid)
}
