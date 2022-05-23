package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//BasicAuth function return echo.MiddlewareFunc for http basic auth
func BasicAuth(config BasicAuthConfig) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == config.Username && password == config.Password {
			return true, nil
		}
		return false, nil
	})
}
