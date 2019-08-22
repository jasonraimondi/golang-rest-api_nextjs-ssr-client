package models

import (
	"fmt"
)

type App struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func NewApp(
	id int64,
	name string,
) *App {
	return &App{
		ID:   id,
		Name: name,
	}
}

func (p *App) GetID() string {
	return fmt.Sprintf("%d", p.ID)
}
