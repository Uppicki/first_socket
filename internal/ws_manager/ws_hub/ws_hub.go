package wshub

import (
	wsclient "first_socket/internal/ws_manager/ws_client"
	wsmessage "first_socket/internal/ws_manager/ws_message"
)

type WSHub interface {
	Disauthorize(string)

	AddClient(wsclient.WSClient)
	RemoveClientByName(string)
	GetClientsWithoutClientName(string) []string

	SendAll()
	SendClientByName(string, wsmessage.WSMessage)
	SendWithoutClientName(string, wsmessage.WSMessage)
}
