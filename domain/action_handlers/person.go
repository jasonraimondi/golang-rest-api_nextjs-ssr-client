package action_handlers

import (
	action2 "git.jasonraimondi.com/jason/jasontest/domain/action"
	model2 "git.jasonraimondi.com/jason/jasontest/domain/model"
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
)

type CreatePersonHandler struct {
	repository.PersonRepository
}

func NewCreatePersonHandler(personRepository repository.PersonRepository) *CreatePersonHandler {
	return &CreatePersonHandler{PersonRepository: personRepository}
}

func (h *CreatePersonHandler) Handle(s *action2.CreatePerson) (err error) {
	p := model2.NewPerson(s.Email)
	if s.First != nil {
		p.FirstName = model2.ToNullString(*s.First)
	}
	if s.Last != nil {
		p.LastName = model2.ToNullString(*s.Last)
	}
	if s.Password != nil {
		if err = p.SetPassword(*s.Password); err != nil {
			return err
		}
	}
	return h.PersonRepository.Create(p)
}
