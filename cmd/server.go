package main

import (
	"log"

	"github.com/goofinator/hasher_nats_server/internal/app"
	"github.com/goofinator/hasher_nats_server/internal/init/startup"
	"github.com/goofinator/hasher_nats_server/internal/web"
)

func main() {
	natsSettings := startup.New()

	session, err := web.IniNats(natsSettings)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	app.Process(session)
}
