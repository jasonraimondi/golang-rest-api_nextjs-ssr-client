package models

import (
	"fmt"
)

type Tag struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func NewTag(
	id int64,
	name string,
) *Tag {
	return &Tag{
		ID:   id,
		Name: name,
	}
}

func (p *Tag) GetID() string {
	return fmt.Sprintf("%d", p.ID)
}
