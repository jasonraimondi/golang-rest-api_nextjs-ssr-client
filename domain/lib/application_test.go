package lib_test

import (
	"fmt"
	"git.jasonraimondi.com/jason/jasontest/domain/action"
	"git.jasonraimondi.com/jason/jasontest/domain/lib"
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
