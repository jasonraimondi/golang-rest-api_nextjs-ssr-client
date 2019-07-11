package lib

import (
	"fmt"
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
	"github.com/jmoiron/sqlx"
)

type Application struct {
	dbx    *sqlx.DB
}

func NewApplication(dbx *sqlx.DB) *Application {
	return &Application{dbx: dbx}
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
