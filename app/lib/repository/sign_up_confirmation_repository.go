package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type SignUpConfirmationRepository struct {
	qb squirrel.StatementBuilderType
	dbx *gorm.DB
}

func (r *SignUpConfirmationRepository) GetByToken(t string) (s models.SignUpConfirmation, err error) {
	token := uuid.FromStringOrNil(t)
	s = models.SignUpConfirmation{}
	err = r.dbx.First(&token, "token = ?", token).Error
	return s, err
}

func (r *SignUpConfirmationRepository) Delete(s *models.SignUpConfirmation) error {
	eq := squirrel.Eq{"token": s.Token}
	sql, args, err := r.qb.Delete("sign_up_confirmation").Where(eq).ToSql()
	if err != nil {
		return err
	}
	return r.dbx.Raw(sql, args...).Error
}
func (r *SignUpConfirmationRepository) Create(s *models.SignUpConfirmation) error {
	return r.dbx.Create(s).Error
}
