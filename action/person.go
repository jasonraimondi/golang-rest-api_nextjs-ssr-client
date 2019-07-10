package action

import (
	"git.jasonraimondi.com/jason/jasontest/lib"
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreatePerson struct {
	*lib.Command
	First    *string
	Last     *string
	Email    string
	Password *string
}

func NewCreatePerson(first *string, last *string, email string, password *string) *CreatePerson {
	command := &lib.Command{
		Time:      time.Now(),
		CommandId: uuid.NewV4().String(),
	}
	return &CreatePerson{command, first, last, email, password}
}

type GetPersonByEmail struct {
	Email string
}

func NewGetPersonByEmail(email string) *GetPersonByEmail {
	return &GetPersonByEmail{email}
}
