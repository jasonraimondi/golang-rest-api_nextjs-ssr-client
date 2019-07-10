package lib

import (
	"fmt"
	"git.jasonraimondi.com/jason/jasontest/repository"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/jmoiron/sqlx"
)

type Application struct {
	dbx    *sqlx.DB
	Driver *database.Driver
}

func NewApplication(dbx *sqlx.DB, driver *database.Driver) *Application {
	return &Application{dbx: dbx, Driver: driver}
}

func (a *Application) RepositoryFactory() *repository.Factory {
	return repository.NewFactory(a.dbx)
}

func (a *Application) Dispatch(command interface{}) (err error) {
	fmt.Println(command)
	//for _, v := range commands {
	// get command handler for command
	// dispatch that command to the handler
	//}
	return err
}
