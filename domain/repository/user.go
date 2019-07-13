package repository

import (
	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	dbx *sqlx.DB
}

func NewUserRepository(dbx *sqlx.DB) *UserRepository {
	return &UserRepository{dbx}
}

var (
	create = `
INSERT INTO users (
	id, 
	first_name, 
	last_name, 
	email, 
	password_hash, 
	is_verified, 
	created_at, 
	modified_at
)
VALUES (
	:id, 
	:first_name, 
	:last_name, 
	:email, 
	:password_hash, 
	:is_verified, 
	:created_at, 
	:modified_at
)
`
	update = `
UPDATE users 
	SET 
		id=:id,
		first_name=:first_name, 
		last_name=:last_name,
		email=:email,
		password_hash=:password_hash,
		is_verified=:is_verified,
		modified_at=:modified_at
WHERE id=$1
	`
)

func (r *UserRepository) GetById(id string) (*model.User, error) {
	user := &model.User{}
	if err := r.dbx.Get(user, `SELECT * FROM users WHERE id=$1`, id); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := r.dbx.Get(user, `SELECT * FROM users WHERE email=$1`, email)
	return user, err
}

func (r *UserRepository) Create(u model.User) (err error) {
	_, err = r.dbx.NamedExec(create, u)
	return err
}

func UpdateUserTx(tx *sqlx.Tx, u model.User) {
	u.ModifiedAt = model.ToNullInt64Now()
	if _, err := tx.NamedExec(update, u); err != nil {
		panic(err)
	}
}

func CreateUserTx(tx *sqlx.Tx, u model.User) (err error) {
	_, err = tx.NamedExec(create, u)
	return err
}

func GetByIdTx(tx *sqlx.Tx, token string) (*model.User, error) {
	user := &model.User{}
	if err := tx.Get(user, `SELECT * FROM users WHERE id=$1`, token); err != nil {
		return nil, err
	}
	return user, nil
}
