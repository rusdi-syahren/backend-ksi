package external

import (
	"github.com/Klinisia/backend-ksi/src/shared"
)

// Notif interface
type Notif interface {
	SendMessage(WhatsappPayload) shared.Output
	CheckNumber(WhatsappPayloadCheckNum) shared.Output
}
