package models

import (
	"fmt"
)

type Tag struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:255"`
	//Photos []PhotoHandler `gorm:"many2many:photo_tag"json:"-"`
}

func (p *Tag) GetID() string {
	return fmt.Sprintf("%d", p.ID)
}
