package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World!"

	assert.Equal(t, got, want)
}
