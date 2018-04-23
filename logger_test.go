package middleware

import (
	"bytes"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/some", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	b := new(bytes.Buffer)

	LoggerWithOutput(b)(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})(c)

	if !strings.Contains(b.String(), "200") {
		t.Errorf("Wrong error message: %s", b.String())
	}
}
