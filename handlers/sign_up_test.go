package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"git.jasonraimondi.com/jason/jasontest/test/utils"
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
	h := utils.NewTestHandler()

	if assert.NoError(t, h.SignUp(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, `{"message":"Created"}
`, rec.Body.String())
	}
}
