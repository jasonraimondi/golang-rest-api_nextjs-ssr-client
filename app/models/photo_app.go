package models

import (
	uuid "github.com/satori/go.uuid"
)

type PhotoApp struct {
	PhotoID uuid.UUID
	Photo   *Photo
	AppID   uint `gorm:"column:tag_id"`
	App     *Tag
}

func (p *PhotoApp) TableName() string {
	return "photo_app"
}
