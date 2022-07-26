package external

import (
	"github.com/Klinisia/backend-ksi/src/shared"
	"github.com/labstack/echo/v4"
)

// Payment interface
type Payment interface {
	Charge(echo.Context, interface{}) shared.Output
	GetToken(echo.Context, TokenParams) shared.Output
	CheckStatusPayment(echo.Context, string) shared.Output
}
