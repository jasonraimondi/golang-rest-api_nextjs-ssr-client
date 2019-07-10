package action

import (
	lib2 "git.jasonraimondi.com/jason/jasontest/domain/lib"
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreatePerson struct {
	*lib2.Command
	First    *string
	Last     *string
	Email    string
	Password *string
}

func NewCreatePerson(first *string, last *string, email string, password *string) *CreatePerson {
	command := &lib2.Command{
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
