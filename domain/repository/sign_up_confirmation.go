package repository

import (
	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	//uuid "github.com/satori/go.uuid"
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

func (r *SqlxSignUpConfirmationRepository) GetByToken(t string) (s *model.SignUpConfirmation, err error) {
	token := uuid.FromStringOrNil(t)
	s = &model.SignUpConfirmation{}
	if err = r.dbx.Get(s, `SELECT * FROM sign_up_confirmation WHERE token=$1`, token); err != nil {
		return nil, err
	}
	return s, nil
}

func DeleteSignUpConfirmationTx(tx *sqlx.Tx, s *model.SignUpConfirmation) {
	tx.MustExec(`DELETE FROM sign_up_confirmation WHERE token=$1`, s.Token)
}

func GetByTokenTx(tx *sqlx.Tx, t string) (s *model.SignUpConfirmation, err error) {
	token := uuid.FromStringOrNil(t)
	s = &model.SignUpConfirmation{}
	if err = tx.Get(s, `SELECT * FROM sign_up_confirmation WHERE token=$1`, token); err != nil {
		return nil, err
	}
	return s, nil
}

func CreateSignUpConfirmationTx(tx *sqlx.Tx, s *model.SignUpConfirmation) {
	tx.MustExec(
		"INSERT INTO sign_up_confirmation (token, user_id, created_at) VALUES ($1, $2, $3)",
		s.Token,
		s.UserId,
		s.CreatedAt,
	)
}
