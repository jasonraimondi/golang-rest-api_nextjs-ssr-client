package web

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConnection(t *testing.T) {
	_, connected := RepositoryFactory()

	assert.True(t, connected)
}
