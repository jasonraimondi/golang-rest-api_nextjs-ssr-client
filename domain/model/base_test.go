package model_test

import (
	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNullString(t *testing.T) {
	invalid := model.ToNullString("")
	valid := model.ToNullString("jason")

	assert.False(t, invalid.Valid)
	assert.True(t, valid.Valid)
}
