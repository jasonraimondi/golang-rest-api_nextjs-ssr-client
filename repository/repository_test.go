package repository_test

import (
	"git.jasonraimondi.com/jason/jasontest/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)



func TestGetConnection(t *testing.T) {
	r, err := repository.Initialize()
	assert.NoError(t, err)

	u := r.Person()
	assert.NoError(t, err)
	assert.NotNil(t, u)

	assert.True(t, true)
	//jason := models.Person{}
	//err = driver.Get(&jason, "SELECT * FROM person WHERE first_name=$1", "Jason")
}
