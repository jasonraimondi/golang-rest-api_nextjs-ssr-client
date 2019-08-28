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
	FileName    string    `gorm:"type:varchar(100); not null"`
	RelativeURL string    `gorm:"type:varchar(255); not null"`
	SHA256      string    `gorm:"type:varchar(64); not null"`
	MimeType    string    `gorm:"type:varchar(100); not null"`
	FileSize    uint64    `gorm:"not null"`
	Apps        []Tag     `gorm:"many2many:photo_app"`
	Tags        []Tag     `gorm:"many2many:photo_tag"`
	UserID      uuid.UUID `gorm:"not null"`
	User        *User
	Description sql.NullString
	Width       sql.NullInt64
	Height      sql.NullInt64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

func NewPhoto(
	uid uuid.UUID,
	u *User,
	fileName string,
	sha256 string,
	mimeType string,
	fileSize uint64,
) *Photo {
	id := uid.String()
	return &Photo{
		ID:          uid,
		UserID:      u.ID,
		FileName:    fileName,
		RelativeURL: fmt.Sprintf("%s/%s%s", id[:2], id, strings.ToLower(filepath.Ext(fileName))),
		SHA256:      sha256,
		MimeType:    mimeType,
		FileSize:    fileSize,
	}
}

func (p *Photo) GetID() string {
	return p.ID.String()
}

func (p *Photo) AddTags(tags []Tag) {
	for _, tag := range tags {
		p.AddTag(tag)
	}
}

func (p *Photo) AddTag(tag Tag) {
	p.Tags = append(p.Tags, tag)
}

func (p *Photo) AddApps(apps []Tag) {
	for _, app := range apps {
		p.AddApp(app)
	}
}

func (p *Photo) AddApp(app Tag) {
	p.Apps = append(p.Apps, app)
}
