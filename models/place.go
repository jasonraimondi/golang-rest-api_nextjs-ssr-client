package models

import (
	"database/sql"
)

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}
