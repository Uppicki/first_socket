package wsservicerepository

import (
	"first_socket/pkg/ws_service/client"

	"github.com/gorilla/websocket"
)

type IClientRepository interface {
	CreateClient(string, string, *websocket.Conn) wsserviceclient.IWSClient
	AddClient(wsserviceclient.IWSClient) error

	RemoveUser(string)

	RemoveClient(string, string)

	GetUserClients(string) []wsserviceclient.IWSClient

	GetUserWithoutClient(string, string) []wsserviceclient.IWSClient

	GetUsersClients([]string) []wsserviceclient.IWSClient
}
