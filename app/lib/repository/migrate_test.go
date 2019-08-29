package repository_test

import (
	"testing"

	"github.com/jinzhu/gorm"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
)

func TestMigrate(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failure creating something")
	}
	defer db.Close()
	repository.Migrate(db)
	//var sql = "INSERT INTO photo_tag (photo_id, tag_id) VALUES (?, ?)"
	//err = db.Raw(sql, uuid.NewV4(), 1).Error
	//if err == nil {
	//	t.Fatalf("missing foreign keys")
	//}
}
