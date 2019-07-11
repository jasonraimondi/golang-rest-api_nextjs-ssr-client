package repository

import (
	model2 "git.jasonraimondi.com/jason/jasontest/domain/model"
	"github.com/jmoiron/sqlx"
)

type PersonRepository interface {
	GetById(id string) (*model2.Person, error)
	GetByEmail(email string) (*model2.Person, error)
	Create(person model2.Person) error
}

type SqlxPersonRepository struct {
	dbx *sqlx.DB
}

func NewSqlxPersonRepository(dbx *sqlx.DB) *SqlxPersonRepository {
	return &SqlxPersonRepository{dbx}
}

func (r *SqlxPersonRepository) GetById(id string) (p model2.Person, err error) {
	p = model2.Person{}
	err = r.dbx.Get(p, `SELECT * FROM persons WHERE id=$1`, id)
	return p, err
}

func (r *SqlxPersonRepository) GetByEmail(email string) (p model2.Person, err error) {
	p = model2.Person{}
	err = r.dbx.Get(p, `SELECT * FROM persons WHERE email=$1`, email)
	return p, err
}

func (r *SqlxPersonRepository) Create(p model2.Person) (err error) {
	_, err = r.dbx.NamedExec(`
		INSERT INTO persons (id, first_name, last_name, email, password_hash, created_at, modified_at) 
		VALUES (:id, :first_name, :last_name, :email, :password_hash, :created_at, :modified_at)
	`, p)
	return err
}
