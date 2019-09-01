package repository

import (
	"github.com/jinzhu/gorm"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

func Migrate(db *gorm.DB) {
	var tables = []interface{}{
		&models.Photo{},
		&models.App{},
		&models.Tag{},
		&models.PhotoTag{},
		&models.User{},
		&models.SignUpConfirmation{},
	}

	db = db.AutoMigrate(tables...)
	// @todo this randomly fails...
	db.Model(&models.Photo{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.PhotoTag{}).AddForeignKey("photo_id", "photos(id)", "CASCADE", "CASCADE")
	db.Model(&models.PhotoTag{}).AddForeignKey("tag_id", "tags(id)", "CASCADE", "CASCADE")

	db.Model(&models.SignUpConfirmation{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
