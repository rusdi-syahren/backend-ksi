package external

import (
	"github.com/Klinisia/backend-ksi/src/shared"
)

// Notif interface
type Sms interface {
	SendSms(WhatsappPayload) shared.Output
}
