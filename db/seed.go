package db

import (
	"git.jasonraimondi.com/jason/jasontest/action"
	"git.jasonraimondi.com/jason/jasontest/lib"
)

type SeedData struct {
	a *lib.Application
}

func (s *SeedData) Seed() error {
	//s.people()
	return nil
}

func (s *SeedData) people() error {
	var commands []interface{}
	commands = append(commands, action.CreatePerson{})
	return s.a.Dispatch(commands)
	return nil
}