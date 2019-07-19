package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"git.jasonraimondi.com/jason/jasontest/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/models"
	"git.jasonraimondi.com/jason/jasontest/test/utils"
)

func TestSignUpConfirmationRepository_GetByToken(t *testing.T) {
	r := utils.NewTestApplication().RepositoryFactory
	user := models.NewUser("jason@raimondi.us")
	confirmation := models.NewSignUpConfirmation(*user)

	tx := r.DBx.MustBegin()
	repository.CreateUserTx(tx, user)
	repository.CreateSignUpConfirmationTx(tx, confirmation)
	assert.NoError(t, tx.Commit())

	sut1, err := r.SignUpConfirmation().GetByToken(confirmation.Token.String())
	if assert.NoError(t, err) {
		assert.Equal(t, confirmation.Token, sut1.Token)
		assert.Equal(t, confirmation.UserId, sut1.UserId)
	}
}
