package models

import (
	"fmt"
)

type Category struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func NewCategory(
	id int64,
	name string,
) *Category {
	return &Category{
		ID:   id,
		Name: name,
	}
}

func (p *Category) GetID() string {
	return fmt.Sprintf("%d", p.ID)
}
