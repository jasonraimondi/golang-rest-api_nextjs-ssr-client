package models

import (
	"fmt"
)

type App struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:255"`
	//Photos []PhotoHandler `json:"-"`
}

func (p *App) GetID() string {
	return fmt.Sprintf("%d", p.ID)
}
