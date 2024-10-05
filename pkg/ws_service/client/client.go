package wsserviceclient

import (
	"first_socket/pkg/ws_service/payload/request"
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

func (client *wsClient) Run() {
	client.mu.Lock()
	client.isActive = true
	client.mu.Unlock()
	go client.readerRun()
	go client.writerRun()
}

func (client *wsClient) readerRun() {
	defer func() {
		client.mu.Lock()
		if client.isActive {
			client.receivedMessage <- wsservicemessage.DisconnectedMessage(
				client.ownerLogin,
			)
			client.isActive = false
		}
		client.mu.Unlock()
	}()

	client.receivedMessage <- wsservicemessage.ConnectedMessage(
		client.ownerLogin,
	)

	for client.isActive {
		var req wsservicerequests.IWSRequest

		if err := client.conn.ReadJSON(&req); err != nil {
			break
		}

		if message, err := req.ToMessage(); err == nil {
			client.receivedMessage <- message
		} else {
			break
		}
	}
}

func (client *wsClient) writerRun() {
	for client.isActive {
		select {
		case message := <-client.sendedMessage:
			response, err := message.ToResponse()
			if err != nil {
				break
			}

			if innerErr := client.conn.WriteJSON(response); innerErr != nil {
				break
			}
		}
	}
}

func (client *wsClient) GetReceivedChan() <-chan wsservicemessage.IWSMessage {
	return client.receivedMessage
}

func (client *wsClient) Send(message wsservicemessage.IWSMessage) {
	select {
	case client.sendedMessage <- message:
	default:
	}
}

func (client *wsClient) Close() {
}

func (client *wsClient) GetOwnerLogin() string {
	return client.ownerLogin
}

func (client *wsClient) GetConnKey() string {
	return client.connKey
}
