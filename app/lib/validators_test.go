package lib_test

import (
	"testing"

	"gopkg.in/go-playground/validator.v9"

	"git.jasonraimondi.com/jason/jasontest/app/lib"
)

func TestValidatorsPass(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("password-strength", lib.ValidatePasswordStrength)
	if err != nil {
		t.Fatalf("error registering password strength validator")
	}
	weakPass := "hi"
	strongPass := "12345678" // to us, strong is > 7 characters, nothing else nothing less

	if err = validate.Var(weakPass, "required,password-strength"); err == nil {
		t.Fatalf("weak password is not failing password-strength validator")
	}
	if err = validate.Var(strongPass, "required,password-strength"); err != nil {
		t.Fatalf("strong pass should be valid password-strength")
	}
}
