package web

import (
	"fmt"

	"github.com/goofinator/hasher_nats_server/internal/api"
	"github.com/goofinator/hasher_nats_server/internal/init/startup"
	"github.com/goofinator/hasher_nats_server/pkg"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

// NatsSession contains the data using for nats communocation
type natsSession struct {
	Incoming     chan pkg.Message
	connection   *nats.EncodedConn
	subscription *nats.Subscription
	natsSettings startup.NatsSettings
}

// IniNats connect to the nats server and recieve all messages into the chanel
func IniNats(natsSettings startup.NatsSettings) (api.NatsSession, error) {
	session := natsSession{natsSettings: natsSettings}

	nc, err := nats.Connect(natsSettings.URL)
	if err != nil {
		return nil, fmt.Errorf("error on Connect: %s", err)
	}
	session.connection, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, fmt.Errorf("error on Connect: %s", err)
	}

	session.Incoming = make(chan pkg.Message)

	session.subscription, err = session.connection.BindRecvChan("worker.*.out", session.Incoming)
	if err != nil {
		return nil, fmt.Errorf("error on Subscribe: %s", err)
	}

	return session, nil
}

func (ns natsSession) UUID() uuid.UUID {
	return ns.natsSettings.UUID
}

// Close closes subscription, drains and closes connection,
// closes Incoming chanel
func (ns natsSession) Close() {
	ns.subscription.Unsubscribe()
	ns.connection.Drain()
	ns.connection.Close()
	close(ns.Incoming)
}

func (ns natsSession) DataSource() <-chan pkg.Message {
	return ns.Incoming
}

// SendMessage sends message via nats
func (ns natsSession) SendMessage(subject string, msg pkg.Message) error {
	if err := ns.connection.Publish(subject, msg); err != nil {
		return fmt.Errorf("error on SendMessage: %s", err)
	}
	return nil
}
