package web

import (
	"database/sql"
	"git.jasonraimondi.com/jason/learn-with-tests/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type RepositoryFactory struct {
	DB *sqlx.DB
}

func (r RepositoryFactory) User() *UserRepository {
	return &UserRepository{db: r.DB}
}

func (r RepositoryFactory) Migrate(schema string) error {
	_ = r.DB.MustExec(schema)
	return nil
}

func (r RepositoryFactory) Seed() (err error) {
	tx := r.DB.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	tx.MustExec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "United States", "New York", "1")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Hong Kong", "852")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Singapore", "65")
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	person := &models.Person{
		FirstName: sql.NullString{"Jane", true},
		LastName: sql.NullString{"Citizen", true},
		Email: "jane.citzen@example.com",
	}
	_, err = tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", person)
	err = tx.Commit()
	return err
}
