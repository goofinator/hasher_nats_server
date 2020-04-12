package startup

import (
	"flag"

	"github.com/google/uuid"
)

// Default values of IniData fields
const (
	DefaultURL        = "nats://localhost:4222"
	DefaultChanelName = "hasher"
)

// IniData structure stores initial data to start a app
type IniData struct {
	URL    string
	Sender uuid.UUID
}

// New returns IniData obtained from user or DefaultPort
func New() *IniData {
	iniData := &IniData{}
	flag.StringVar(&iniData.URL, "url", DefaultURL, "url of nats service")
	flag.Parse()

	iniData.Sender = uuid.New()

	return iniData
}
