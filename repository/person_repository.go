package repository

import (
	"git.jasonraimondi.com/jason/jasontest/models"
	"github.com/jmoiron/sqlx"
)

type PersonRepository struct {
	db *sqlx.DB
}

func (r PersonRepository) GetById(id string) (p *models.Person, err error) {
	p = &models.Person{}
	err = r.db.Get(p, "SELECT * FROM person WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return p, nil
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
