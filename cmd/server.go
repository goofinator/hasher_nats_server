package main

import (
	"github.com/goofinator/hasher_nats_server/internal/app"
	"github.com/goofinator/hasher_nats_server/internal/init/startup"
	"github.com/goofinator/hasher_nats_server/internal/web"
)

func main() {
	iniData := startup.New()
	session := web.IniNats(iniData)
	defer session.Close()

	app.Process(session, iniData)
}
