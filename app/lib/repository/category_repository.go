package repository

import (
	"github.com/jmoiron/sqlx"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type AppRepository struct {
	dbx *sqlx.DB
}

var createApp = `
	INSERT INTO apps (id, name)
	VALUES (:id, :name)
`

func (r *AppRepository) GetById(id string) (tag *models.App, err error) {
	tag = &models.App{}
	if err = r.dbx.Get(tag, `SELECT * FROM apps WHERE id=$1`, id); err != nil {
		return nil, err
	}
	return tag, nil
}

func CreateAppTx(tx *sqlx.Tx, u *models.App) (err error) {
	_, err = tx.NamedExec(createApp, u)
	return err
}
