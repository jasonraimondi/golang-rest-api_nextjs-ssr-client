package repository

import (
	"github.com/jmoiron/sqlx"

	"git.jasonraimondi.com/jason/jasontest/models"
)

type PhotoRepository struct {
	dbx *sqlx.DB
}

var createPhoto = `
	INSERT INTO photos (id, file_name, relative_url, mime_type, sha256, file_size, width, height, user_id, created_at, modified_at)
	VALUES (:id, :file_name, :relative_url, :mime_type, :sha256, :file_size, :width, :height, :user_id, :created_at, :modified_at)
`

func (r *PhotoRepository) GetById(id string) (photo *models.Photo, err error) {
	photo = &models.Photo{}
	if err = r.dbx.Get(photo, `SELECT * FROM photos WHERE id=$1`, id); err != nil {
		return nil, err
	}
	return photo, nil
}

func CreatePhotoTx(tx *sqlx.Tx, u *models.Photo) (err error) {
	_, err = tx.NamedExec(createPhoto, u)
	return err
}
