package repository

import (
	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetById(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(p model.User) error
}

type SqlxUserRepository struct {
	dbx *sqlx.DB
}

func NewSqlxUserRepository(dbx *sqlx.DB) *SqlxUserRepository {
	return &SqlxUserRepository{dbx}
}

func (r *SqlxUserRepository) GetById(id string) (p model.User, err error) {
	p = model.User{}
	err = r.dbx.Get(p, `SELECT * FROM users WHERE id=$1`, id)
	return p, err
}

func (r *SqlxUserRepository) GetByEmail(email string) (p *model.User, err error) {
	p = &model.User{}
	err = r.dbx.Get(p, `SELECT * FROM users WHERE email=$1`, email)
	return p, err
}

var insertQuery = `
	INSERT INTO users (id, first_name, last_name, email, password_hash, is_verified, created_at, modified_at) 
	VALUES (:id, :first_name, :last_name, :email, :password_hash, :is_verified, :created_at, :modified_at)
`

func (r *SqlxUserRepository) Create(p *model.User) (err error) {
	_, err = r.dbx.NamedExec(insertQuery, p)
	return err
}

func (r *SqlxUserRepository) CreateTx(tx *sqlx.Tx, p *model.User) (err error) {
	_, err = tx.NamedExec(insertQuery, p)
	return err
}