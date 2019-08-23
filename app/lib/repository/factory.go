package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Factory struct {
	dbx *sqlx.DB
	qb  squirrel.StatementBuilderType
}

func NewFactory(dbx *sqlx.DB) *Factory {
	return &Factory{
		dbx,
		squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *Factory) QueryBuilder() squirrel.StatementBuilderType {
	return r.qb
}

func (r *Factory) DB() *sqlx.DB {
	return r.dbx
}

func (r *Factory) User() *UserRepository {
	return &UserRepository{r.DB()}
}

func (r *Factory) SignUpConfirmation() *SignUpConfirmationRepository {
	return &SignUpConfirmationRepository{r.DB()}
}

func (r *Factory) PhotoRepository() *PhotoRepository {
	return &PhotoRepository{r.DB()}
}

func (r *Factory) AppRepository() *AppRepository {
	return &AppRepository{r.qb, r.dbx}
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

func (r *Factory) ListAppsRepository() *ListAppsRepository {
	return &ListAppsRepository{r.qb, r.dbx}
}
