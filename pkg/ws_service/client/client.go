package wsserviceclient

import (
	"first_socket/pkg/ws_service/ws_message"
	"sync"

	"github.com/gorilla/websocket"
)

type wsClient struct {
	conn            *websocket.Conn
	ownerLogin      string
	connKey         string
	sendedMessage   chan wsservicemessage.IWSMessage
	receivedMessage chan wsservicemessage.IWSMessage
	isActive        bool
	mu              sync.Mutex
}

func (client *wsClient) Run() {}

func (client *wsClient) GetReceivedChan() <-chan wsservicemessage.IWSMessage {
	return client.receivedMessage
}

func NewWSClient() IWSClient {
	return &wsClient{}
}
