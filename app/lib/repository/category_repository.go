package repository

import (
	"github.com/jmoiron/sqlx"

	"git.jasonraimondi.com/jason/jasontest/app/models"
)

type CategoryRepository struct {
	dbx *sqlx.DB
}

var createCategory = `
	INSERT INTO categories (id, name)
	VALUES (:id, :name)
`

func (r *CategoryRepository) GetById(id string) (tag *models.Category, err error) {
	tag = &models.Category{}
	if err = r.dbx.Get(tag, `SELECT * FROM categories WHERE id=$1`, id); err != nil {
		return nil, err
	}
	return tag, nil
}

func CreateCategoryTx(tx *sqlx.Tx, u *models.Category) (err error) {
	_, err = tx.NamedExec(createCategory, u)
	return err
}
