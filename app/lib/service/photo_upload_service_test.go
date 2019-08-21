package service_test

import (
	"os"
	"testing"

	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
)

func TestGetFileSHA256(t *testing.T) {
	expected := "6ffd6e5978fece70f17fce35a0346322c101fd6db74495aef33ed8d762775ea1"
	dat, err := os.Open("/tmp/adobegc.log")
	if err != nil {
		t.Errorf("error opening file")
	}
	sha, err := service.GetFileSHA256(dat)
	if err != nil {
		t.Errorf("error opening file")
	}
	if sha != expected {
		t.Errorf("actual: %v != expected: %v", sha, expected)
	}
}
