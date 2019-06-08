package actions

import (
	"git.jasonraimondi.com/jason/jasontest/models"
	"git.jasonraimondi.com/jason/jasontest/repository"
)

type CreatePerson struct {
	first *string
	last *string
	email string
	password *string
}

type CreatePersonHandler struct {
	personRepository *repository.PersonRepository
}

func (h *CreatePersonHandler) Handle(s *CreatePerson) (err error) {
	p := models.NewPerson(s.email)
	if s.first != nil {
		p.FirstName = models.ToNullString(*s.first)
	}
	if s.last != nil {
		p.LastName = models.ToNullString(*s.last)
	}
	if s.password != nil {
		if err = p.SetPassword(*s.password); err != nil {
			return err
		}
	}
	return h.personRepository.Create(p)
}
