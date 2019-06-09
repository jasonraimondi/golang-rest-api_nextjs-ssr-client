package repository

import (
	"git.jasonraimondi.com/jason/jasontest/model"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	dbx *sqlx.DB
}

func (r *Person) GetById(id string) (p *model.Person, err error) {
	p = &model.Person{}
	err = r.dbx.Get(p, `SELECT * FROM persons WHERE id=$1`, id)
	return p, err
}

func (r *Person) GetByEmail(email string) (p *model.Person, err error) {
	p = &model.Person{}
	err = r.dbx.Get(p, `SELECT * FROM persons WHERE email=$1`, email)
	return p, err
}

func (r *Person) Create(p model.Person) (err error) {
	_, err = r.dbx.NamedExec(`
		INSERT INTO persons (id, first_name, last_name, email, password_hash, created_at, modified_at) 
		VALUES (:id, :first_name, :last_name, :email, :password_hash, :created_at, :modified_at)
	`, p)
	return err
}
