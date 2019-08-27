package repository

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) GetById(id string) (*models.User, error) {
	user := &models.User{}
	if uid, err := uuid.FromString(id); err != nil {
		return nil, err
	} else if err = r.db.First(&user, "id = ?", uid).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (user models.User, err error) {
	user = models.User{}
	err = r.db.First(&user, "email = ?", email).Error
	return user, err
}

func (r *UserRepository) Update(u models.User) (err error) {
	return r.db.Save(&u).Error
}

func (r *UserRepository) Create(u models.User) (err error) {
	return r.db.Create(&u).Error
}
