package action

import (
	"git.jasonraimondi.com/jason/jasontest/model"
	"git.jasonraimondi.com/jason/jasontest/repository"
)

type CreatePerson struct {
	First    *string
	Last     *string
	Email    string
	Password *string
}

type CreatePersonHandler struct {
	PersonRepository *repository.Person
}

func (h *CreatePersonHandler) Handle(s *CreatePerson) (err error) {
	p := model.NewPerson(s.Email)
	if s.First != nil {
		p.FirstName = model.ToNullString(*s.First)
	}
	if s.Last != nil {
		p.LastName = model.ToNullString(*s.Last)
	}
	if s.Password != nil {
		if err = p.SetPassword(*s.Password); err != nil {
			return err
		}
	}
	return h.PersonRepository.Create(p)
}
