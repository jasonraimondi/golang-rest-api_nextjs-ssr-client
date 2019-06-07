package repository

import (
	"git.jasonraimondi.com/jason/jasontest/models"
	"github.com/jmoiron/sqlx"
)

type PersonRepository struct {
	db *sqlx.DB
}

func (r PersonRepository) Create(p models.Person) (err error) {
	tx := r.db.MustBegin()
	_, err = tx.NamedExec(`
		INSERT INTO person (id, first_name, last_name, email, created_at, modified_at) 
		VALUES (:id, :first_name, :last_name, :email, created_at, modified_at)
	`, p)
	if err != nil {
		return err
	}
	return nil
}
