package startup

import (
	"flag"

	"github.com/google/uuid"
)

// Default values of NatsSettings fields
const (
	DefaultURL = "nats://localhost:4222"
)

// NatsSettings structure stores initial data to start a app
type NatsSettings struct {
	URL  string
	UUID uuid.UUID
}

// New returns NatsSettings obtained from user or DefaultPort
func New() *NatsSettings {
	natsSettings := &NatsSettings{}
	flag.StringVar(&natsSettings.URL, "url", DefaultURL, "url of nats service")
	flag.Parse()

	natsSettings.UUID = uuid.New()

	return natsSettings
}
