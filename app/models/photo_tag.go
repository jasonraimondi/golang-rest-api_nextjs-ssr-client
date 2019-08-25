package models

type PhotoTag struct {
	Photo *Photo
	Tag   *Tag
}

func (PhotoTag) TableName() string {
	return "photo_tag"
}
