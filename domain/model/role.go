package model

type Role struct {
	ID   int32  `db:"id"`
	Name string `db:"name"`
}
