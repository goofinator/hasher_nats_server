package remotes

import (
	"log"

	"github.com/goofinator/hasher_nats_server/internal/init/startup"
	"github.com/nats-io/nats.go"
)

// IniNats connect to the nats server and recieve all messages into the chanel
func IniNats(iniData *startup.IniData) (*nats.Conn, chan *nats.Msg) {
	nc, err := nats.Connect(iniData.URL)
	if err != nil {
		log.Fatalf("error on Connect: %s", err)

	}

	ch := make(chan *nats.Msg)

	nc.ChanSubscribe(iniData.ChanelName, ch)
	return nc, ch
}
