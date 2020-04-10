package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

// Process performs main cycle
func Process(ch <-chan *nats.Msg) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

EXIT:
	for {
		select {
		case msg := <-ch:
			go processMessage(msg)
		case <-sigs:
			break EXIT
		}
	}
}

func processMessage(msg *nats.Msg) {
	fmt.Printf("message geted: %s\n", msg.Data)
}
