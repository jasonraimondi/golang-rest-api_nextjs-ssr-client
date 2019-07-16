package model

import (
	"database/sql"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Photo struct {
	ID           uuid.UUID     `db:"id"`
	OriginalName string        `db:"original_name"`
	ContentType  string        `db:"content_type"`
	FileSize     int64         `db:"file_size"`
	UserId       uuid.UUID     `db:"user_id"`
	CreatedAt    int64         `db:"created_at"`
	ModifiedAt   sql.NullInt64 `db:"modified_at"`
}

func NewPhoto(id uuid.UUID, u *User, originalName string, contentType string, fileSize int64) *Photo {
	return &Photo{
		ID:           id,
		UserId: 	  u.ID,
		OriginalName: originalName,
		ContentType:  contentType,
		FileSize:     fileSize,
		ModifiedAt:   ToNullInt64(""),
		CreatedAt:    time.Now().Unix(),
	}
}
