package external

import (
	"github.com/labstack/echo/v4"
	"github.com/rusdi-syahren/backend-ksi/src/shared"
)

// Payment interface
type Payment interface {
	Charge(echo.Context, interface{}) shared.Output
	GetToken(echo.Context, TokenParams) shared.Output
	CheckStatusPayment(echo.Context, string) shared.Output
}
