package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/goofinator/hasher_nats_server/internal/api"
	"github.com/goofinator/hasher_nats_server/internal/init/startup"
	"github.com/goofinator/hasher_nats_server/internal/utils"
	"github.com/goofinator/hasher_nats_server/pkg"
	"github.com/google/uuid"
)

// Process performs main cycle
func Process(session api.NatsSession, iniData *startup.IniData) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

EXIT:
	for {
		select {
		case msg := <-session.DataSource():
			go ProcessMessage(session, iniData, msg)
		case <-sigs:
			break EXIT
		}
	}
}

// ProcessMessage processes message and send response via session
func ProcessMessage(session api.NatsSession, iniData *startup.IniData, msg *pkg.Message) {
	hashes := utils.Hash(msg.Body)
	result, err := json.Marshal(hashes)
	if err != nil {
		result = make([]byte, 0)
	}

	reply := &pkg.Message{
		Sender: iniData.Sender,
		ID:     uuid.New(),
		Type:   pkg.DefaultMessageType,
		Body:   result,
	}

	subject := fmt.Sprintf("worker.%s.in", msg.Sender)
	if err := session.SendMessage(subject, reply); err != nil {
		log.Printf("error on SendMessage: %s\n", err)
	}
}
