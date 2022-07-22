package external

import (
	"github.com/Klinisia/backend-ksi/src/shared"
	"github.com/labstack/echo"
)

// Shipment interface
type Shipment interface {
	GetShipmentCost(echo.Context, *ShipmentCost) shared.Output
	GetProvice(echo.Context, *ShipmentCost) (interface{}, error)
	GetCity(echo.Context, *ShipmentCost) (interface{}, error)
}
