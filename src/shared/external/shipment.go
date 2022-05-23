package external

import (
	"github.com/labstack/echo"
	"gitlab.com/k1476/scaffolding/src/shared"
)

// Shipment interface
type Shipment interface {
	GetShipmentCost(echo.Context, *ShipmentCost) shared.Output
	GetProvice(echo.Context, *ShipmentCost) (interface{}, error)
	GetCity(echo.Context, *ShipmentCost) (interface{}, error)
}
