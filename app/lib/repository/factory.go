package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/gorm"
)

type Factory struct {
	dbx *gorm.DB
	qb  squirrel.StatementBuilderType
}

func NewFactory(dbx *gorm.DB) *Factory {
	qb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)
	return &Factory{dbx, qb}
}

func (r *Factory) QueryBuilder() squirrel.StatementBuilderType {
	return r.qb
}

func (r *Factory) DB() *gorm.DB {
	return r.dbx
}

func (r *Factory) User() *UserRepository {
	return &UserRepository{r.dbx}
}

func (r *Factory) SignUpConfirmation() *SignUpConfirmationRepository {
	return &SignUpConfirmationRepository{r.qb, r.dbx}
}

func (r *Factory) PhotoRepository() *PhotoRepository {
	return &PhotoRepository{r.dbx}
}

func (r *Factory) TagRepository() *TagRepository {
	return &TagRepository{r.qb, r.dbx}
}

func (r *Factory) ListPhotosRepository() *ListPhotosRepository {
	return &ListPhotosRepository{r.qb, r.dbx}
}

func (r *Factory) ListTagsRepository() *ListTagsRepository {
	return &ListTagsRepository{r.qb, r.dbx}
}
//
//func (r *Factory) ListAppsRepository() *ListAppsRepository {
//	return &ListAppsRepository{r.qb, r.dbx}
//}
