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
	URL        string
	ChanelName string
	Sender     uuid.UUID
}

// Configuration returns port to use obtained from user or DefaultPort
func Configuration() *IniData {
	iniData := &IniData{}
	flag.StringVar(&iniData.URL, "url", DefaultURL, "url of nats service")
	flag.StringVar(&iniData.ChanelName, "chanel", DefaultChanelName, "chanel to use")
	flag.Parse()

	iniData.Sender = uuid.New()

	return iniData
}
