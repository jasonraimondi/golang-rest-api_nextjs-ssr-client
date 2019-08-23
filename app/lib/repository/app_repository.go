package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type AppRepository struct {
	queryBuilder squirrel.StatementBuilderType
	dbx          *sqlx.DB
}

func (r *AppRepository) Delete(id string) error {
	sql, args, err := r.queryBuilder.Delete("apps").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}
	r.dbx.MustExec(sql, args...)
	return nil
}

func (r *AppRepository) UnlinkFromPhoto(photoId string, appId int64) error {
	sql, args, err := r.queryBuilder.
		Delete("photo_app").
		Where(squirrel.Eq{
			"app_id":   appId,
			"photo_id": photoId,
		}).ToSql()
	if err != nil {
		return err
	}
	r.dbx.MustExec(sql, args...)
	return nil
}
