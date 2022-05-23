package external

import (
	"github.com/labstack/echo"
	"gitlab.com/k1476/scaffolding/src/shared"
)

// Payment interface
type Payment interface {
	Charge(echo.Context, interface{}) shared.Output
	GetToken(echo.Context, TokenParams) shared.Output
	CheckStatusPayment(echo.Context, string) shared.Output
}
