//go:generate mockgen -destination=./mocks/mock_nat_session.go -package=mocks github.com/goofinator/hasher_nats_server/internal/api NatsSession
//go:generate goimports -w ./mocks/mock_nat_session.go

package api

import "github.com/goofinator/hasher_nats_server/pkg"

// NatsSession wraps methods to manipulate nats session
type NatsSession interface {
	DataSource() <-chan *pkg.Message
	Close()
	SendMessage(subject string, msg *pkg.Message) error
}
