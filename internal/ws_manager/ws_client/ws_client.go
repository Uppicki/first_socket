package wsclient

import wsmessage "first_socket/internal/ws_manager/ws_message"

type WSClient interface {
	GetNameOwnerUser() string
	Run()
	Close()
	Send(wsmessage.WSMessage)
	GetReceivedChan() <-chan wsmessage.WSMessage
}
