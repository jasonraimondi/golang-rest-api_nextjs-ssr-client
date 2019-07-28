package repository

import (
	"git.jasonraimondi.com/jason/jasontest/models"
	"github.com/jmoiron/sqlx"
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

func (r *PhotoRepository) CountForUser(userId string) (count int64, err error) {
	rows, err := r.dbx.Query("SELECT COUNT(*) as count FROM photos WHERE user_id=$1", userId)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		if err:= rows.Scan(&count); err != nil {
			return 0, err
		}
	}
	return count, nil
}

func (r *PhotoRepository) ListForUser(userId string, page int64, itemsPerPage int64) ([]models.Photo, error) {

	limit := itemsPerPage
	offset := limit * (page - 1)

	photos := []models.Photo{}
	err := r.dbx.Select(&photos, `SELECT * FROM photos WHERE user_id=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, userId, limit, offset)
	if err != nil {
		return nil, err
	}
	return photos, nil
}

func CreatePhotoTx(tx *sqlx.Tx, u *models.Photo) (err error) {
	_, err = tx.NamedExec(createPhoto, u)
	return err
}
