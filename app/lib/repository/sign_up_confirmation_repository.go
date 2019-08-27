package repository

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type SignUpConfirmationRepository struct {
	db *gorm.DB
}

func (r *SignUpConfirmationRepository) GetByToken(t string) (s models.SignUpConfirmation, err error) {
	token := uuid.FromStringOrNil(t)
	s = models.SignUpConfirmation{}
	err = r.db.First(&s, "token = ?", token).Error
	return s, err
}

func (r *SignUpConfirmationRepository) Delete(s *models.SignUpConfirmation) error {
	return r.db.Delete(s, "token = ?", s.Token).Error
}
func (r *SignUpConfirmationRepository) Create(s *models.SignUpConfirmation) error {
	return r.db.Create(s).Error
}
