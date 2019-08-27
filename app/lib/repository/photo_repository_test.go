package repository_test

import (
	"testing"

	"git.jasonraimondi.com/jason/jasontest/app/models"
	"git.jasonraimondi.com/jason/jasontest/app/test/utils"
)

func TestPhotoRepository_ForTags(t *testing.T) {
	tables := []interface{}{
		&models.Photo{},
		&models.Tag{},
	}
	a := utils.NewTestApplication(tables)
	pr := a.RepositoryFactory.PhotoRepository()
	tags := []string{"tag-one", "two", "tres"}
	paginator := pr.ForTags(tags, 1, 25)
	if paginator.TotalRecord != 0 {
		t.Fatalf("how did you get a record?")
	}
}