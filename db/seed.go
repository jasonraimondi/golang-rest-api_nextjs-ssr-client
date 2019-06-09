package db

import (
	"github.com/jmoiron/sqlx"
)

type SeedData struct {
	dbx *sqlx.DB
}

func (s *SeedData) Seed() error {
	s.people()
	return nil
}

func (s *SeedData) people() {

}