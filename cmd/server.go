package main

import (
	"github.com/goofinator/hasher_nats_server/internal/app"
	"github.com/goofinator/hasher_nats_server/internal/init/startup"
	"github.com/goofinator/hasher_nats_server/internal/remotes"
	"github.com/goofinator/hasher_nats_server/pkg"
	"github.com/nats-io/nats.go"
)

func main() {
	iniData := startup.Configuration()
	session := remotes.IniNats(iniData)
	defer session.Close()

	app.Process(session, iniData)
}

func cleanup(connection *nats.EncodedConn, ch chan<- *pkg.Message) {
	connection.Drain()
	connection.Close()
	close(ch)
}
