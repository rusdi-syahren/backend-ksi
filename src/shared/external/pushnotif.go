package external

import "context"

// Firebase abstract interface
type Firebase interface {
	SendNotification(ctx context.Context, payload *Payload) <-chan []byte
}
