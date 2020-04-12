package main

import (
	"github.com/goofinator/hasher_nats_server/internal/app"
	"github.com/goofinator/hasher_nats_server/internal/init/startup"
	"github.com/goofinator/hasher_nats_server/internal/web"
)

func main() {
	natsSettings := startup.New()
	session := web.IniNats(natsSettings)
	defer session.Close()

	app.Process(session, natsSettings)
}
