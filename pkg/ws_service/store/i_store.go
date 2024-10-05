package wsservicestore

import (
	"first_socket/pkg/ws_service/client"
)

type IStore interface {
	AddClient(wsserviceclient.IWSClient) error

	RemoveUser(string)
	RemoveClient(string, string)

	GetUserClients(string) []wsserviceclient.IWSClient
	GetUserWithoutClient(string, string) []wsserviceclient.IWSClient

	GetUsersClients([]string) []wsserviceclient.IWSClient
}
