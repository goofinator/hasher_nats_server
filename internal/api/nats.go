//go:generate mockgen -destination=./mocks/mock_nat_session.go -package=mocks github.com/goofinator/hasher_nats_server/internal/api NatsSession
//go:generate goimports -w ./mocks/mock_nat_session.go

package api

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

// NatsSession wraps methods to manipulate nats session
type NatsSession interface {
	DataSource() <-chan *Message
	Close()
	SendMessage(subject string, msg *Message) error
}
