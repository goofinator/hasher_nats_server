package main

import (
	"github.com/goofinator/hasher_nats_server/internal/app"
	"github.com/goofinator/hasher_nats_server/internal/init/startup"
	"github.com/goofinator/hasher_nats_server/internal/remotes"
	"github.com/nats-io/nats.go"
)

func main() {
	iniData := startup.Configuration()
	nc, ch := remotes.IniNats(iniData)
	defer cleanup(nc, ch)

	app.Process(ch)
}

func cleanup(nc *nats.Conn, ch chan<- *nats.Msg) {
	nc.Drain()
	nc.Close()
	close(ch)
}
