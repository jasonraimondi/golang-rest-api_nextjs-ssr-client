package service_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
	"git.jasonraimondi.com/jason/jasontest/app/models"
	"git.jasonraimondi.com/jason/jasontest/app/test/utils"
)

func TestTagService_AddTagsToPhoto(t *testing.T) {
	tables := []interface{}{
		&models.Photo{},
		&models.Tag{},
	}
	a := utils.NewTestApplication(tables)

	user := models.NewUser("jason@raimondi.us")
	photo := models.NewPhoto(uuid.NewV4(), user, "myfilename.png", "mysha256", "image/png", 42)
	photo.AddTag(models.Tag{Name: "og-tag"})
	photo.AddTag(models.Tag{Name: "alpha"})

	if err := a.RepositoryFactory.PhotoRepository().Create(photo); err != nil {
		t.Fatalf("error creating photo")
	}

	sut, _ := a.RepositoryFactory.PhotoRepository().GetById(photo.GetID())
	newTags := []string{"alpha", "beta", "zeta"}
	if err := a.ServiceFactory.TagService().AddTagsToPhoto(sut, newTags); err != nil {
		t.Fatalf("error adding tags to photo")
	}

	sut, err := a.RepositoryFactory.PhotoRepository().GetById(sut.GetID())
	if err != nil {
		t.Fatalf("error fetching photo")
	}
	if len(sut.Tags) != 4 {
		t.Fatalf("should have 4 tags (%d)", len(sut.Tags))
	} else if tag := sut.Tags[0].Name; tag != "og-tag" {
		t.Fatalf("expected: %s actual: %s", "og-tag", tag)
	} else if tag := sut.Tags[1].Name; tag != "alpha" {
		t.Fatalf("expected: %s actual: %s", "alpha", tag)
	} else if tag := sut.Tags[2].Name; tag != "beta" {
		t.Fatalf("expected: %s actual: %s", "beta", tag)
	} else if tag := sut.Tags[3].Name; tag != "zeta" {
		t.Fatalf("expected: %s actual: %s", "zeta", tag)
	}
}

func TestDifference(t *testing.T) {
	a := []string{"a", "b", "c", "d", "extra"}
	b := []string{"b", "c", "d", "funny"}
	sut := service.ArrayDiff(a, b)
	if sut[0] != "a" {
		t.Fatalf("invalid array diff")
	}
	if sut[1] != "extra" {
		t.Fatalf("invalid array diff")
	}
}
