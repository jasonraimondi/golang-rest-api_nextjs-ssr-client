package repository

import (
	"git.jasonraimondi.com/jason/jasontest/models"
	"github.com/jmoiron/sqlx"
)

type PersonRepository struct {
	dbx *sqlx.DB
}

func (r *PersonRepository) GetById(id string) (p *models.Person, err error) {
	p = &models.Person{}
	err = r.dbx.Get(p, `SELECT * FROM persons WHERE id=$1`, id)
	return p, err
}

func (r *PersonRepository) GetByEmail(email string) (p *models.Person, err error) {
	p = &models.Person{}
	err = r.dbx.Get(p, `SELECT * FROM persons WHERE email=$1`, email)
	return p, err
}

func (r *PersonRepository) Create(p models.Person) (err error) {
	_, err = r.dbx.NamedExec(`
		INSERT INTO persons (id, first_name, last_name, email, password_hash, created_at, modified_at) 
		VALUES (:id, :first_name, :last_name, :email, :password_hash, :created_at, :modified_at)
	`, p)
	return err
}
