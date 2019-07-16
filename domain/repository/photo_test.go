package repository_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"git.jasonraimondi.com/jason/jasontest/domain/lib"
	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
)

func TestPhotoRepository_GetById(t *testing.T) {
	r := lib.NewTestApplication().RepositoryFactory
	user := model.NewUser("jason@raimondi.us")
	photos := model.NewPhoto(uuid.NewV4(), user, "filename", "image title", 213241)

	tx := r.DBx.MustBegin()
	assert.NoError(t, repository.CreateUserTx(tx, user))
	assert.NoError(t, repository.CreatePhotoTx(tx, photos))
	assert.NoError(t, tx.Commit())

	sut1, err := r.PhotoRepository().GetById(photos.ID.String())
	if assert.NoError(t, err) {
		assert.Equal(t, photos.FileSize, sut1.FileSize)
		assert.Equal(t, photos.OriginalName, sut1.OriginalName)
	}
}
