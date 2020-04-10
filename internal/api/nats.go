package api

import "github.com/gofrs/uuid"

// Messge describes nats message
type Messge struct {
	Type   uint8
	Body   []byte
	ID     uuid.UUID
	Sender uuid.UUID
}
