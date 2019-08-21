package service_test

import (
	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/models"
	"git.jasonraimondi.com/jason/jasontest/app/test/utils"

	"testing"

	"github.com/stretchr/testify/assert"
)

func xTestService_ValidateEmailSignUpConfirmation(t *testing.T) {
	a := utils.NewTestApplication()
	user := models.NewUser("jason@raimondi.us")
	confirmation := models.NewSignUpConfirmation(*user)

	tx := a.RepositoryFactory.DB().MustBegin()
	assert.NoError(t, repository.CreateUserTx(tx, user))
	repository.CreateSignUpConfirmationTx(tx, confirmation)
	assert.NoError(t, tx.Commit())

	err := a.ServiceFactory.SignUpService().ValidateEmailSignUpConfirmation(confirmation.Token.String(), confirmation.UserId.String())
	if assert.NoError(t, err) {

	}
}
