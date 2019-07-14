package service_test

import (
	"git.jasonraimondi.com/jason/jasontest/domain/lib"
	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func xTestService_ValidateEmailSignUpConfirmation(t *testing.T) {
	a := lib.NewTestApplication()
	user := model.NewUser("jason@raimondi.us")
	confirmation := model.NewSignUpConfirmation(*user)

	tx := a.RepositoryFactory.DBx.MustBegin()
	assert.NoError(t, repository.CreateUserTx(tx, user))
	repository.CreateSignUpConfirmationTx(tx, confirmation)
	assert.NoError(t, tx.Commit())

	err := a.ServiceFactory.ValidateEmailSignUpConfirmation(confirmation.Token.String(), confirmation.UserId.String())
	if assert.NoError(t, err) {

	}
}
