package service_test

import (
	"testing"

	"git.jasonraimondi.com/jason/jasontest/app/models"
	"git.jasonraimondi.com/jason/jasontest/app/test/utils"
)

func TestService_ValidateEmailSignUpConfirmation(t *testing.T) {
	tables := []interface{}{
		&models.User{},
		&models.SignUpConfirmation{},
	}
	a := utils.NewTestApplication(tables)

	user := models.NewUser("jason@raimondi.us")
	if err := a.RepositoryFactory.User().Create(*user); err != nil {
		t.Fatalf("error creating user")
	}
	confirmation := models.NewSignUpConfirmation(*user)
	if err := a.RepositoryFactory.SignUpConfirmation().Create(confirmation); err != nil {
		t.Fatalf("error creating sign_up_confirmation")
	}
	err := a.ServiceFactory.SignUpService().ValidateEmailSignUpConfirmation(confirmation.Token.String(), confirmation.UserID.String())
	if err != nil {
		t.Fatalf("error validating sign_up_confirmation")
	}
}
