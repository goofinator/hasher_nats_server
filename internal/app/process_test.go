package app_test

import (
	"bytes"
	"errors"
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
	"github.com/stretchr/testify/assert"
)

const (
	testDataFile = "process_test.txt"
)

var (
	clientUUID   = uuid.New()
	serverUUID   = uuid.New()
	natsSettings = &startup.NatsSettings{
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

func TestProcessMessageSuccess(t *testing.T) {
	//catch log
	logBuf := &bytes.Buffer{}
	log.SetOutput(logBuf)

	//You need a NatsSession mock to process the request
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	ns := mocks.NewMockNatsSession(mockController)

	setProcessMessageExpectations(t, ns, nil)

	ProcessMessage(ns, natsSettings, message)
	assert.Empty(t, logBuf.String())
}

func TestProcessMessageFail(t *testing.T) {
	//catch log
	logBuf := &bytes.Buffer{}
	log.SetOutput(logBuf)

	//You need a NatsSession mock to process the request
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	ns := mocks.NewMockNatsSession(mockController)

	setProcessMessageExpectations(t, ns, errors.New("some error"))

	ProcessMessage(ns, natsSettings, message)
	assert.Contains(t, logBuf.String(), string("error on SendMessage: some error\n"))
}

func setProcessMessageExpectations(t *testing.T, ns *mocks.MockNatsSession, err error) {
	subject := fmt.Sprintf("worker.%s.in", clientUUID)

	ns.EXPECT().
		SendMessage(subject, NewMatcherMessage(t, message)).
		Return(err)
}

func NewMatcherMessage(t *testing.T, msg *pkg.Message) gomock.Matcher {
	reply := &pkg.Message{
		ID:   msg.ID,
		Body: getTestBody(),
	}
	return &matcherMessage{
		msg: reply,
		t:   t,
	}
}

// some fields of reply are randomly generated
// so we can't compare it strictly with something
// and should provide some matcher
type matcherMessage struct {
	msg *pkg.Message
	t   *testing.T
}

func (m *matcherMessage) Matches(x interface{}) bool {
	msg, ok := x.(*pkg.Message)
	if !ok {
		m.t.Errorf("fail to x.(*pkg.Message) in Matches")
		return false
	}
	switch {
	case msg.Sender != serverUUID:
		m.t.Errorf("msg.Sender != serverUUID")
	case msg.Type != pkg.DefaultMessageType:
		m.t.Errorf("msg.Type != pkg.DefaultMessageType")
	case bytes.Compare(msg.Body, m.msg.Body) != 0:
		m.t.Errorf("msg body is not equal to specified in %q file", testDataFile)
	case msg.ID == m.msg.ID || msg.ID == clientUUID || msg.ID == serverUUID:
		m.t.Errorf("msg.ID should differ from clientUUID, serverUUID, recieved message ID")
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
