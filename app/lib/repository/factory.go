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

func (r *Factory) ListPhotosRepository() *ListPhotosRepository {
	return &ListPhotosRepository{
		queryBuilder: r.qb,
		dbx:          r.DB(),
	}
}

func (r *Factory) ListTagsRepository() *ListTagsRepository {
	return &ListTagsRepository{
		queryBuilder: r.qb,
		dbx:          r.DB(),
	}
}
