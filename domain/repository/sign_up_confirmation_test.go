package repository_test

import (
	"git.jasonraimondi.com/jason/jasontest/domain/lib"
	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignUpConfirmationRepository_GetByToken(t *testing.T) {
	r := lib.NewTestApplication().RepositoryFactory
	user := model.NewUser("jason@raimondi.us")
	confirmation := model.NewSignUpConfirmation(*user)

	tx := r.DBx.MustBegin()
	repository.CreateUserTx(tx, *user)
	repository.CreateSignUpConfirmationTx(tx, confirmation)
	assert.NoError(t, tx.Commit())

	sut1, err := r.SignUpConfirmation().GetByToken(confirmation.Token.String())
	if assert.NoError(t, err) {
		assert.Equal(t, confirmation.Token, sut1.Token)
		assert.Equal(t, confirmation.UserId, sut1.UserId)
	}
}