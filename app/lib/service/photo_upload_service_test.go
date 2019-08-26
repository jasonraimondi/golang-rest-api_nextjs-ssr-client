package service_test

import (
	"os"
	"testing"

	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
)

func TestGetFileSHA256(t *testing.T) {
	expected := "bc23f77592655381c789040b1d3d5c0c8cb82b510cd87fcac9c90d09336dd343"
	dat, err := os.Open("../../test/attachments/codecraft-2018.jpg")
	if err != nil {
		t.Fatalf("error opening file")
	}
	sha, err := service.GetFileSHA256(dat)
	if err != nil {
		t.Fatalf("error getting file sha256")
	}
	if sha != expected {
		t.Fatalf("actual: %v != expected: %v", sha, expected)
	}
}
