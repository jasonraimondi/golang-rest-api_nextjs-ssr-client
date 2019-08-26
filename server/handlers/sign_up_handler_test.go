package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"git.jasonraimondi.com/jason/jasontest/app/models"
	"git.jasonraimondi.com/jason/jasontest/app/test/utils"

	"github.com/labstack/echo"
)

func TestCreateUser(t *testing.T) {
	e := echo.New()
	f := make(url.Values)
	f.Set("first", "Jon")
	f.Set("last", "Snow")
	f.Set("email", "jon23@example.com")
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	tables := []interface{}{
		&models.User{},
		&models.SignUpConfirmation{},
	}
	h := utils.NewTestHandler(tables)
	if err := h.SignUpHandler().SignUp(c); err != nil {
		t.Fatalf("error signing up")
	}
	if rec.Code != http.StatusCreated {
		t.Fatalf("invalid status code")
	}

	content := strings.Trim(rec.Body.String(), "\n")
	if content != `{"message":"Created"}` {
		t.Fatalf("invalid body content (%s)", content)
	}
}
