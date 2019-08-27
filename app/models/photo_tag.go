package models

import (
	uuid "github.com/satori/go.uuid"
)

type PhotoTag struct {
	PhotoID uuid.UUID
	Photo   *Photo
	TagID   uint
	Tag     *Tag
}

func (p *PhotoTag) TableName() string {
	return "photo_tag"
}