package repository_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"git.jasonraimondi.com/jason/jasontest/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/models"
	"git.jasonraimondi.com/jason/jasontest/test/utils"
)

func TestPhotoRepository_GetById(t *testing.T) {
	r := utils.NewTestApplication().RepositoryFactory
	user := models.NewUser("jason@raimondi.us")
	photos := models.NewPhoto(uuid.NewV4(), user, "filename", "image title", "", 213241)

	tx := r.DBx.MustBegin()
	assert.NoError(t, repository.CreateUserTx(tx, user))
	assert.NoError(t, repository.CreatePhotoTx(tx, photos))
	assert.NoError(t, tx.Commit())

	sut1, err := r.PhotoRepository().GetById(photos.GetID())
	if assert.NoError(t, err) {
		assert.Equal(t, photos.FileSize, sut1.FileSize)
		assert.Equal(t, photos.FileName, sut1.FileName)
	}
}
