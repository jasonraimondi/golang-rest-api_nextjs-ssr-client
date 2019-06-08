package repository_test

import (
	"git.jasonraimondi.com/jason/jasontest/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitialize(t *testing.T) {
	_, err := repository.NewTestDB()
	assert.NoError(t, err)
}
