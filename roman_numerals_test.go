package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var romanNumeral = []struct {
	given    string // input
	expected int    // expected result
}{
	{"I", 1},
	{"II", 2},
	{"III", 3},
}

func TestRomanNumeral(t *testing.T) {
	for _, tt := range romanNumeral {
		actual := RomanNumeral(tt.given)
		assert.Equal(t, actual, tt.expected)
	}
}
