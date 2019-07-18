package repository

import (
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/models"
)

type SignUpConfirmationRepository struct {
	dbx *sqlx.DB
}

func (r *SignUpConfirmationRepository) GetByToken(t string) (s *models.SignUpConfirmation, err error) {
	token := uuid.FromStringOrNil(t)
	s = &models.SignUpConfirmation{}
	if err = r.dbx.Get(s, `SELECT * FROM sign_up_confirmation WHERE token=$1`, token); err != nil {
		return nil, err
	}
	return s, nil
}

func DeleteSignUpConfirmationTx(tx *sqlx.Tx, s *models.SignUpConfirmation) {
	tx.MustExec(`DELETE FROM sign_up_confirmation WHERE token=$1`, s.Token)
}

func GetByTokenTx(tx *sqlx.Tx, t string) (s *models.SignUpConfirmation, err error) {
	token := uuid.FromStringOrNil(t)
	s = &models.SignUpConfirmation{}
	if err = tx.Get(s, `SELECT * FROM sign_up_confirmation WHERE token=$1`, token); err != nil {
		return nil, err
	}
	return s, nil
}

func CreateSignUpConfirmationTx(tx *sqlx.Tx, s *models.SignUpConfirmation) {
	tx.MustExec(
		"INSERT INTO sign_up_confirmation (token, user_id, created_at) VALUES ($1, $2, $3)",
		s.Token,
		s.UserId,
		s.CreatedAt,
	)
}
