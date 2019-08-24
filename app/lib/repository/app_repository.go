package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/gorm"
)

type AppRepository struct {
	qb  squirrel.StatementBuilderType
	dbx *gorm.DB
}

func (r *AppRepository) Delete(id string) error {
	eq := squirrel.Eq{"id": id}
	sql, args, err := r.qb.Delete("apps").Where(eq).ToSql()
	if err != nil {
		return err
	}
	return r.dbx.Raw(sql, args...).Error
}

func (r *AppRepository) UnlinkFromPhoto(photoId string, appId int64) error {
	eq := squirrel.Eq{"app_id": appId, "photo_id": photoId}
	sql, args, err := r.qb.Delete("photo_app").Where(eq).ToSql()
	if err != nil {
		return err
	}
	return r.dbx.Raw(sql, args...).Error
}
