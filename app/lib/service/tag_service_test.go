package service_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/models"
	"git.jasonraimondi.com/jason/jasontest/app/test/utils"
)

func TestPhotoAppService_AddTagsToPhoto(t *testing.T) {
	tables := []interface{}{
		&models.Photo{},
		&models.Tag{},
	}
	a := utils.NewTestApplication(tables)

	user := models.NewUser("jason@raimondi.us")
	photo := models.NewPhoto(uuid.NewV4(), user, "myfilename", "mysha256", "image/png", 42)
	photo.AddTag(models.Tag{Name: "og-tag"})
	photo.AddTag(models.Tag{Name: "alpha"})

	if err := a.RepositoryFactory.PhotoRepository().Create(photo); err != nil {
		t.Fatalf("error creating photo")
	}

	newTags := []string{"alpha", "beta", "zeta"}
	if err := a.ServiceFactory.PhotoAppService().AddTagsToPhoto(photo.GetID(), newTags); err != nil {
		t.Fatalf("error adding tags to photo")
	}

	photo, err := a.RepositoryFactory.PhotoRepository().GetById(photo.GetID())
	if err != nil {
		t.Fatalf("error fetching photo")
	}
	if len(photo.Tags) != 4 {
		t.Fatalf("should have three tags")
	} else if photo.Tags[0].Name != "og-tag" {
		t.Fatalf("invalid name (%s)", photo.Tags[0].Name)
	} else if photo.Tags[1].Name != "alpha" {
		t.Fatalf("invalid name (%s)", photo.Tags[0].Name)
	} else if photo.Tags[2].Name != "beta" {
		t.Fatalf("invalid name (%s)", photo.Tags[1].Name)
	} else if photo.Tags[3].Name != "zeta" {
		t.Fatalf("invalid name (%s)", photo.Tags[2].Name)
	}
}