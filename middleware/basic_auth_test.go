package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestBasicAuth(t *testing.T) {

	validUsernameAndPassword := `Basic Ymhpbm5la2E6ZGExYzI1ZDgtMzdjOC00MWIxLWFmZTItNDJkZDQ4MjViZmVh`

	t.Run("Should No Error with Valid Auth", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, validUsernameAndPassword)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "klinisia.id"))
		})

		var config BasicAuthConfig
		config.Username = "klinisia"
		config.Password = "da1c25d8-37c8-41b1-afe2-42dd4825bfea"

		mw := BasicAuth(config)(handler)
		err := mw(c)
		assert.NoError(t, err)
	})

	t.Run("Should Error with Invalid Auth", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "invalidAuth")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "klinisia.id"))
		})

		var config BasicAuthConfig
		config.Username = "klinisia"
		config.Password = "da1c25d8-37c8-41b1-afe2-42dd4825bfea"

		mw := BasicAuth(config)(handler)
		err := mw(c)
		assert.Error(t, err)
	})
}
