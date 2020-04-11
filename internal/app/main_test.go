package app_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/hasher_nats_server/internal/api/mocks"
	. "github.com/goofinator/hasher_nats_server/internal/app"
	"github.com/goofinator/hasher_nats_server/internal/init/startup"
	"github.com/goofinator/hasher_nats_server/pkg"
	"github.com/google/uuid"
)

const (
	testDataFile = "main_test.txt"
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
		ID:   msg.ID,
		Body: getTestBody(),
	}
	return &matcherMessage{msg: reply}
}

// some fields of reply are randomly generated
// so we can't compare it strictly with something
// and should provide some matcher
type matcherMessage struct {
	msg *pkg.Message
}

func (m *matcherMessage) Matches(x interface{}) bool {
	msg, ok := x.(*pkg.Message)
	switch {
	case !ok:
		log.Fatal("fail to x.(*pkg.Message) in Matches")
	case msg.Sender != serverUUID:
		log.Println("msg.Sender != serverUUID")
	case msg.Type != pkg.DefaultMessageType:
		log.Println("msg.Type != pkg.DefaultMessageType")
	case bytes.Compare(msg.Body, m.msg.Body) != 0:
		log.Printf("msg body is not equal to specified in %q file", testDataFile)
	case msg.ID == m.msg.ID || msg.ID == clientUUID || msg.ID == serverUUID:
		log.Println("msg.ID should differ from clientUUID, serverUUID, recieved message ID")
	default:
		return true
	}
	return false
}

func (m *matcherMessage) String() string {
	return fmt.Sprintf("is a %v", m.msg)
}

func getTestBody() []byte {
	tb, err := ioutil.ReadFile(testDataFile)
	if err != nil {
		log.Fatalf("can't open %q file", testDataFile)
	}
	result := make([]byte, 0)
	fmt.Sscanf(string(tb), "%X", &result)
	return result
}
