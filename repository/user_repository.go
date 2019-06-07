package repository

import (
	"git.jasonraimondi.com/jason/jasontest/models"
	"github.com/jmoiron/sqlx"
)

type PersonRepository struct {
	db *sqlx.DB
}

func (r PersonRepository) create(u models.Person) (bool, error) {
	tx := r.db.MustBegin()
	_, err := tx.NamedExec("INSERT INTO person (id, first_name, last_name, email) VALUES (:id, :first_name, :last_name, :email)", u)
	if err != nil {
		return false, err
	}
	return true, nil
}
