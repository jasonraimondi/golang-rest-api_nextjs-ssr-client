package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

func ConnectToSQLiteInMemory() (driver *sqlx.DB) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic("error connecting to sqlite3 :memory:")
	}
	return sqlx.NewDb(db, "sqlite3")
}
