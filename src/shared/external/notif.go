package external

import (
	"gitlab.com/k1476/scaffolding/src/shared"
)

// Notif interface
type Notif interface {
	SendMessage(WhatsappPayload) shared.Output
	CheckNumber(WhatsappPayloadCheckNum) shared.Output
}
