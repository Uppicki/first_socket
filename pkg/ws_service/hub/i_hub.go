package wsservicehub

import (
	wsserviceclient "first_socket/pkg/ws_service/client"

	"github.com/gorilla/websocket"
)

type IWSClientHub interface {
	AddClient(string, string, *websocket.Conn) (
		wsserviceclient.IWSClient,
		error,
	)
}
