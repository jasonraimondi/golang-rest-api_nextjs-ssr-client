package models_test

import (
	"testing"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

func TestToNullString(t *testing.T) {
	invalid := models.ToNullString("")
	valid := models.ToNullString("jason")
	if invalid.Valid == true {
		t.Fatalf("blank string is not valid")
	}
	if valid.Valid == false {
		t.Fatalf("string is valid")
	}
}

func TestToNullInt64(t *testing.T) {
	invalid := models.ToNullInt64("")
	valid := models.ToNullInt64("634")
	if invalid.Valid == true {
		t.Fatalf("blank int is not valid")
	}
	if valid.Valid == false {
		t.Fatalf("int is valid")
	}
}
