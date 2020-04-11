package pkg

import "github.com/google/uuid"

const (
	// DefaultMessageType is a message Type used by this client
	DefaultMessageType = 0
)

// Message describes nats message
type Message struct {
	Type   uint8
	Body   []byte
	ID     uuid.UUID
	Sender uuid.UUID
}
