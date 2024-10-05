package wsserviceclient

import (
	"first_socket/pkg/ws_service/ws_message"

	"github.com/gorilla/websocket"
)

type IWSClient interface {
	Run()
	Close()

	GetReceivedChan() <-chan wsservicemessage.IWSMessage
	Send(wsservicemessage.IWSMessage)

	GetOwnerLogin() string
	GetConnKey() string
}

func NewWSClient(
	ownerLogin string,
	connKey string,
	conn *websocket.Conn,
) IWSClient {
	return &wsClient{
		ownerLogin:      ownerLogin,
		connKey:         connKey,
		conn:            conn,
		sendedMessage:   make(chan wsservicemessage.IWSMessage),
		receivedMessage: make(chan wsservicemessage.IWSMessage),
		isActive:        false,
	}
}
