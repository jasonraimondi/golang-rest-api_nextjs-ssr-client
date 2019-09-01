package service_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/models"
	"git.jasonraimondi.com/jason/jasontest/app/test/utils"
)

func TestPhotoService_UpdatePhoto(t *testing.T) {
	tables := []interface{}{
		&models.Photo{},
		&models.App{},
		&models.Tag{},
	}
	a := utils.NewTestApplication(tables)

	appName := "jsonsapp"

	user := models.NewUser("jason@raimondi.us")
	photo := models.NewPhoto(uuid.NewV4(), user, "myfilename.png", "mysha256", "image/png", 42)
	photo.App = &models.App{Name: "Reddit"}

	if err := a.RepositoryFactory.PhotoRepository().Create(photo); err != nil {
		t.Fatalf("error creating photo")
	}
	if err := a.ServiceFactory.PhotoService().UpdatePhoto(photo.GetID(), "here is my new description", appName, []string{}); err != nil {
		t.Fatalf("error updating photo")
	}

	photo1, err := a.RepositoryFactory.PhotoRepository().GetById(photo.GetID())
	if err != nil {
		t.Fatalf("error fetching photo")
	}
	if sut := photo1.App.Name; sut != appName {
		t.Fatalf("app name should be %s, got (%s)", appName, sut)
	}
}
