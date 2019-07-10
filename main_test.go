package main_test

import (
	"fmt"
	"git.jasonraimondi.com/jason/jasontest/action"
	"git.jasonraimondi.com/jason/jasontest/lib"
	"testing"
)

func TestApplication(t *testing.T) {
	first := "first"
	last := "last"
	email := "email"
	password := "password"
	c := action.NewCreatePerson(
		&first,
		&last,
		email,
		&password,
	)
	fmt.Println(c)

	a := lib.NewTestApplication()
	_ = a.Dispatch(c)
}
