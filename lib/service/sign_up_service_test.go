package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"git.jasonraimondi.com/jason/jasontest/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/models"
	"git.jasonraimondi.com/jason/jasontest/test/utils"
)

func xTestService_ValidateEmailSignUpConfirmation(t *testing.T) {
	a := utils.NewTestApplication()
	user := models.NewUser("jason@raimondi.us")
	confirmation := models.NewSignUpConfirmation(*user)

	tx := a.RepositoryFactory.DBx.MustBegin()
	assert.NoError(t, repository.CreateUserTx(tx, user))
	repository.CreateSignUpConfirmationTx(tx, confirmation)
	assert.NoError(t, tx.Commit())

	err := a.ServiceFactory.SignUpService().ValidateEmailSignUpConfirmation(confirmation.Token.String(), confirmation.UserId.String())
	if assert.NoError(t, err) {

	}
}
