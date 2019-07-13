package web_test

import (
	"git.jasonraimondi.com/jason/jasontest/web"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	e := echo.New()
	f := make(url.Values)
	f.Set("first_name", "Jon")
	f.Set("last_name", "Snow")
	f.Set("email", "jon23@example.com")
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := web.NewTestHandler()

	if assert.NoError(t, h.SignUp(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "\"Created\"\n", rec.Body.String())
	}
}