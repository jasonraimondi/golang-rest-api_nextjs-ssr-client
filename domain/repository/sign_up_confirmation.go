package repository

import (
	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"github.com/jmoiron/sqlx"
)

type SignUpConfirmationRepository interface {
	GetById(id string) (*model.SignUpConfirmation, error)
	GetByEmail(email string) (*model.SignUpConfirmation, error)
	Create(p model.SignUpConfirmation) error
}

type SqlxSignUpConfirmationRepository struct {
	dbx *sqlx.DB
}

func NewSqlxSignUpConfirmationRepository(dbx *sqlx.DB) *SqlxSignUpConfirmationRepository {
	return &SqlxSignUpConfirmationRepository{dbx}
}

func (r *SqlxSignUpConfirmationRepository) GetById(id string) (p *model.SignUpConfirmation, err error) {
	p = &model.SignUpConfirmation{}
	err = r.dbx.Get(p, `SELECT * FROM users WHERE id=$1`, id)
	return p, err
}

func (r *SqlxSignUpConfirmationRepository) GetByEmail(email string) (p *model.SignUpConfirmation, err error) {
	p = &model.SignUpConfirmation{}
	err = r.dbx.Get(p, `SELECT * FROM users WHERE email=$1`, email)
	return p, err
}

