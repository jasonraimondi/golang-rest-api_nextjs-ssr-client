package repository_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/models"
	"git.jasonraimondi.com/jason/jasontest/app/test/utils"
)

func TestPhotoRepository_ForUser(t *testing.T) {
	tables := []interface{}{
		&models.Photo{},
		&models.User{},
		&models.Tag{},
		&models.App{},
	}
	a := utils.NewTestApplication(tables)
	pr := a.RepositoryFactory.PhotoRepository()

	user := models.NewUser("jason@raimondi.us")
	a.RepositoryFactory.UserRepository().Create(*user)
	p1 := models.NewPhoto(uuid.NewV4(), user, "filename1", "sha1", "image/png", 1234)
	a.RepositoryFactory.PhotoRepository().Create(p1)
	p2 := models.NewPhoto(uuid.NewV4(), user, "filename2", "sha1", "image/png", 1234)
	a.RepositoryFactory.PhotoRepository().Create(p2)
	p3 := models.NewPhoto(uuid.NewV4(), user, "filename3", "sha1", "image/png", 1234)
	a.RepositoryFactory.PhotoRepository().Create(p3)

	user2 := models.NewUser("jason@raimondi.us")
	a.RepositoryFactory.UserRepository().Create(*user2)
	models.NewPhoto(uuid.NewV4(), user2, "filename1", "sha1", "image/png", 1234)

	paginator := pr.ForUser(user.GetID(), 1, 25)
	if paginator.TotalRecord != 3 {
		t.Fatalf("invalid record count")
	}
	var photos *[]models.Photo
	photos = paginator.Records.(*[]models.Photo)
	for k, v := range *photos {
		if name := v.FileName; name != "filename3" && k == 0 {
			t.Fatalf("actual: %s expected: %s", name, "filename1")
		} else if name := v.FileName; name != "filename2" && k == 1 {
			t.Fatalf("actual: %s expected: %s", name, "filename2")
		}
	}
}

//func TestPhotoRepository_ForTags(t *testing.T) {
//	tables := []interface{}{
//		&models.PhotoHandler{},
//		&models.Tag{},
//	}
//	a := utils.NewTestApplication(tables)
//	pr := a.RepositoryFactory.PhotoRepository()
//	tags := []string{"tag-one", "two", "tres"}
//	paginator := pr.ForTags(tags, 1, 25)
//	if paginator.TotalRecord != 0 {
//		t.Fatalf("should have 4 photos (%d)", len(photo.Tags))
//	}
//}
