package models

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Photo struct {
	ID          uuid.UUID `gorm:"primary_key"`
	FileName    string
	RelativeURL string
	SHA256      string
	MimeType    string
	FileSize    int64
	Width       sql.NullInt64
	Height      sql.NullInt64
	UserID      uuid.UUID
	Tags        []*Tag `gorm:"many2many:photo_tag"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

//func (p *Photo) BeforeCreate(scope *gorm.Scope) error {
//	return scope.SetColumn("ID", uuid.NewV4().String())
//}

func NewPhoto(
	id uuid.UUID,
	u *User,
	fileName string,
	sha256 string,
	mimeType string,
	fileSize int64,
) *Photo {
	s := id.String()
	return &Photo{
		ID:          id,
		UserID:      u.ID,
		FileName:    fileName,
		RelativeURL: fmt.Sprintf("%s/%s%s", s[:2], s, strings.ToLower(filepath.Ext(fileName))),
		SHA256:      sha256,
		MimeType:    mimeType,
		FileSize:    fileSize,
	}
}

func (p *Photo) AddTag(tag Tag) {
	p.Tags = append(p.Tags, &tag)
}

func (p *Photo) GetID() string {
	return p.ID.String()
}
