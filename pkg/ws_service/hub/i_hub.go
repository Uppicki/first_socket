package wsservicehub

import (
	"first_socket/pkg/ws_service/client"
	wsservicemessage "first_socket/pkg/ws_service/ws_message"

	"github.com/gorilla/websocket"
)

type IWSClientHub interface {
	AddClient(string, string, *websocket.Conn) (
		wsserviceclient.IWSClient,
		error,
	)
	RemoveUser(string)
	RemoveUserClient(string, string)

	SendUser(string, wsservicemessage.IWSMessage)
	SendUserWithoutClient(string, string, wsservicemessage.IWSMessage)
	SendUsers([]string, wsservicemessage.IWSMessage)
}
