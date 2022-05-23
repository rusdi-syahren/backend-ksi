package external

import (
	"github.com/labstack/echo"
	"gitlab.com/k1476/scaffolding/src/shared"
)

// ChatEngine interface
type ChatEngine interface {
	GetToken(echo.Context, *TokenRequest) shared.Output
	CreateRoom(echo.Context, *RoomRequest) shared.Output
	GetListRoom(echo.Context, *RoomRequest) shared.Output
	SendMessage(*MessageRequest) shared.Output
}
