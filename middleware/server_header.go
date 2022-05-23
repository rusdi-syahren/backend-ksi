package middleware

import (
	"github.com/labstack/echo"
)

// ServerHeader function
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "my-service/1.0")
		return next(c)
	}
}
