package action

import (
	"git.jasonraimondi.com/jason/jasontest/lib"
)

type CreatePerson struct {
	*lib.Command
	First    *string
	Last     *string
	Email    string
	Password *string
}

func NewCreatePerson(command *lib.Command, first *string, last *string, email string, password *string) *CreatePerson {
	return &CreatePerson{Command: command, First: first, Last: last, Email: email, Password: password}
}
