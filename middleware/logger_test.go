package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := echo.HandlerFunc(func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.String(http.StatusOK, "my-service klinisia.id"))
	})

	mw := Logger(handler)
	err := mw(c)
	assert.NoError(t, err)

}
