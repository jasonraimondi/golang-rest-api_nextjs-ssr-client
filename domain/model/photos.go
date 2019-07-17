package model

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Photo struct {
	ID          uuid.UUID     `db:"id"`
	FileName    string        `db:"file_name"`
	RelativeURL string        `db:"relative_url"`
	SHA256      string        `db:"sha256"`
	MimeType    string        `db:"mime_type"`
	FileSize    int64         `db:"file_size"`
	Width       sql.NullInt64 `db:"width"`
	Height      sql.NullInt64 `db:"height"`
	UserId      uuid.UUID     `db:"user_id"`
	CreatedAt   int64         `db:"created_at"`
	ModifiedAt  sql.NullInt64 `db:"modified_at"`
}

func NewPhoto(
	id uuid.UUID,
	u *User,
	fileName string,
	sha256 string,
	mimeType string,
	fileSize int64,
) *Photo {
	return &Photo{
		ID:          id,
		UserId:      u.ID,
		FileName:    fileName,
		RelativeURL: fmt.Sprintf("pictures/%s%s", id.String(), strings.ToLower(filepath.Ext(fileName))),
		SHA256:      sha256,
		MimeType:    mimeType,
		FileSize:    fileSize,
		ModifiedAt:  ToNullInt64(""),
		CreatedAt:   time.Now().Unix(),
	}
}

func (p *Photo) GetID() string {
	return p.ID.String()
}
