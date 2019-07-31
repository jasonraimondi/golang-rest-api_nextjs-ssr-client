package repository

import (
	"git.jasonraimondi.com/jason/jasontest/app/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	dbx *sqlx.DB
}

var (
	createUser = `
				INSERT INTO users (id, first, last, email, password_hash, is_verified, created_at, modified_at)
				VALUES (:id, :first, :last, :email, :password_hash, :is_verified, :created_at, :modified_at)
	`
	updateUser = `
				UPDATE users 
				SET 
					id=:id,
					first=:first, 
					last=:last,
					email=:email,
					password_hash=:password_hash,
					is_verified=:is_verified,
					modified_at=:modified_at
				WHERE id=$1
	`
)

func (r *UserRepository) GetById(id string) (*models.User, error) {
	user := &models.User{}
	if err := r.dbx.Get(user, `SELECT * FROM users WHERE id=$1`, id); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := r.dbx.Get(user, `SELECT * FROM users WHERE email=$1`, email)
	return user, err
}

func (r *UserRepository) Create(u *models.User) (err error) {
	_, err = r.dbx.NamedExec(createUser, u)
	return err
}

func UpdateUserTx(tx *sqlx.Tx, u *models.User) {
	u.ModifiedAt = models.ToNullInt64Now()
	if _, err := tx.NamedExec(updateUser, u); err != nil {
		panic(err)
	}
}

func CreateUserTx(tx *sqlx.Tx, u *models.User) (err error) {
	_, err = tx.NamedExec(createUser, u)
	return err
}

func GetByIdTx(tx *sqlx.Tx, token string) (*models.User, error) {
	user := &models.User{}
	if err := tx.Get(user, `SELECT * FROM users WHERE id=$1`, token); err != nil {
		return nil, err
	}
	return user, nil
}
