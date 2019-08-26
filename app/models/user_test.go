package models_test

import (
	"testing"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

func TestUser_SetPassword(t *testing.T) {
	p := models.NewUser("jason@raimondi.us")
	password := "jasonraimondi"
	if err := p.SetPassword(password); err != nil {
		t.Fatalf("error setting password")
	}

	if p.CheckPassword(password) != true {
		t.Fatalf("error validating password")
	}
}

func TestUser_GetFullName(t *testing.T) {
	p := models.NewUser("jason1@raimondi.us")
	p.SetFirst("Jason")
	p.SetLast("Raimondi")

	p2 := models.NewUser("jason2@raimondi.us")
	p2.SetFirst("Jason")

	p3 := models.NewUser("jason3@raimondi.us")
	p3.SetLast("Raimondi")

	p4 := models.NewUser("jason4@raimondi.us")

	if p.GetFullName() != "Jason Raimondi" {
		t.Fatalf("invalid full name")
	}
	if p2.GetFullName() != "Jason" {
		t.Fatalf("invalid full name")
	}
	if p3.GetFullName() != "Raimondi" {
		t.Fatalf("invalid full name")
	}
	if p4.GetFullName() != "" {
		t.Fatalf("invalid full name")
	}
}

func TestUser_GetFullIdentifier(t *testing.T) {
	p := models.NewUser("jason@raimondi.us")
	p.SetFirst("Jason")
	p.SetLast("Raimondi")

	if p.GetFullIdentifier() != "Jason Raimondi <jason@raimondi.us>" {
		t.Fatalf("invalid full identifier")
	}
}
