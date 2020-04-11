package app_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/hasher_nats_server/internal/api/mocks"
	. "github.com/goofinator/hasher_nats_server/internal/app"
	"github.com/goofinator/hasher_nats_server/internal/init/startup"
	"github.com/goofinator/hasher_nats_server/pkg"
	"github.com/google/uuid"
)

var (
	clientUUID = uuid.New()
	serverUUID = uuid.New()
	iniData    = &startup.IniData{
		Sender: serverUUID,
	}
)
var (
	message = &pkg.Message{
		Sender: clientUUID,
		ID:     uuid.New(),
		Type:   pkg.DefaultMessageType,
		Body:   []byte("message 1"),
	}
	replBody = []uint8{
		0x5b, 0x22, 0x55, 0x45, 0x78, 0x56, 0x52, 0x77, 0x3d, 0x3d, 0x22,
		0x2c, 0x22, 0x63, 0x47, 0x78, 0x31, 0x5a, 0x77, 0x3d, 0x3d, 0x22, 0x5d}
)

func TestProcessMessage(t *testing.T) {
	//You need a NatsSession mock to process the request
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	ns := mocks.NewMockNatsSession(mockController)

	setProcessMessageExpectations(ns, nil)

	ProcessMessage(ns, iniData, message)
}

func setProcessMessageExpectations(ns *mocks.MockNatsSession, err error) {
	subject := fmt.Sprintf("worker.%s.in", clientUUID)

	ns.EXPECT().
		SendMessage(subject, NewMatcherMessage(message)).
		Return(err)
}

func NewMatcherMessage(msg *pkg.Message) gomock.Matcher {
	reply := &pkg.Message{
		ID: msg.ID,
	}
	return &matcherMessage{msg: reply}
}

type matcherMessage struct {
	msg *pkg.Message
}

func (m *matcherMessage) Matches(x interface{}) bool {
	msg, ok := x.(*pkg.Message)
	if !ok {
		return false
	}
	if msg.Sender != serverUUID {
		return false
	}
	if msg.Type != pkg.DefaultMessageType {
		return false
	}
	if !bytes.Equal(msg.Body, replBody) {
		return false
	}
	if msg.ID == m.msg.ID || msg.ID == clientUUID || msg.ID == serverUUID {
		return false
	}
	return true
}

func (m *matcherMessage) String() string {
	return fmt.Sprintf("is a %v", m.msg)
}
