//go:generate mockgen -destination=./mocks/mock_nat_session.go -package=mocks github.com/goofinator/hasher_nats_server/internal/api NatsSession
//go:generate goimports -w ./mocks/mock_nat_session.go

package api

import (
	"github.com/goofinator/hasher_nats_server/pkg"
	"github.com/google/uuid"
)

// NatsSession wraps methods to manipulate nats session
type NatsSession interface {
	Close()
	DataSource() <-chan *pkg.Message
	UUID() uuid.UUID
	SendMessage(subject string, msg *pkg.Message) error
}
