package external

import (
	"github.com/Klinisia/backend-ksi/src/shared"
	"github.com/labstack/echo/v4"
)

// ChatEngine interface
type ChatEngine interface {
	GetToken(echo.Context, *TokenRequest) shared.Output
	CreateRoom(echo.Context, *RoomRequest) shared.Output
	GetListRoom(echo.Context, *RoomRequest) shared.Output
	SendMessage(*MessageRequest) shared.Output
}
