package external

import (
	"github.com/labstack/echo/v4"
	"github.com/rusdi-syahren/backend-ksi/src/shared"
)

// Shipment interface
type Shipment interface {
	GetShipmentCost(echo.Context, *ShipmentCost) shared.Output
	GetProvice(echo.Context, *ShipmentCost) (interface{}, error)
	GetCity(echo.Context, *ShipmentCost) (interface{}, error)
}
