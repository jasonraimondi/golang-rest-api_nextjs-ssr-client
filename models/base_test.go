package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNullString(t *testing.T) {
	isNullString := ToNullString("")
	isValidString := ToNullString("jason")
	assert.False(t, isNullString.Valid)
	assert.True(t, isValidString.Valid)
}
