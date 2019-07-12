package model

import (
	"database/sql"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

//ToNullTime invalidates a pq.ToNullTime if empty, validates if not empty
func ToNullTime(t time.Time) pq.NullTime {
	return pq.NullTime{Time: t, Valid: true}
}

func ToNullNullTime() pq.NullTime {
	return pq.NullTime{Time: time.Now(), Valid: false}
}

//ToNullString invalidates a sql.NullString if empty, validates if not empty
func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

//ToNullInt64 validates a sql.NullInt64 if incoming string evaluates to an integer, invalidates if it does not
func ToNullInt64(s string) sql.NullInt64 {
	i, err := strconv.Atoi(s)
	return sql.NullInt64{Int64: int64(i), Valid: err == nil}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
