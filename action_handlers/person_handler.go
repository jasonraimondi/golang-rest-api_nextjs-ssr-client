package action_handlers

import (
	"git.jasonraimondi.com/jason/jasontest/action"
	"git.jasonraimondi.com/jason/jasontest/model"
	"git.jasonraimondi.com/jason/jasontest/repository"
)

type CreatePersonHandler struct {
	repository.PersonRepository
}

func NewCreatePersonHandler(personRepository repository.PersonRepository) *CreatePersonHandler {
	return &CreatePersonHandler{PersonRepository: personRepository}
}

func (h *CreatePersonHandler) Handle(s *action.CreatePerson) (err error) {
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
