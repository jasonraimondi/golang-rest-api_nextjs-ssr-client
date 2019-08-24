package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name   string
	Photos []*Photo `gorm:"many2many:photo_tag"`
}

func (p *Tag) GetID() string {
	return fmt.Sprintf("%d", p.ID)
}
