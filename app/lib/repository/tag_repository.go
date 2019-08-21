package repository

import (
	"github.com/jmoiron/sqlx"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type TagRepository struct {
	dbx *sqlx.DB
}

var createTag = `
	INSERT INTO tags (id, name)
	VALUES (:id, :name)
`

func (r *TagRepository) GetById(id string) (tag *models.Tag, err error) {
	tag = &models.Tag{}
	if err = r.dbx.Get(tag, `SELECT * FROM tags WHERE id=$1`, id); err != nil {
		return nil, err
	}
	return tag, nil
}

func CreateTagTx(tx *sqlx.Tx, u *models.Tag) (err error) {
	_, err = tx.NamedExec(createTag, u)
	return err
}
