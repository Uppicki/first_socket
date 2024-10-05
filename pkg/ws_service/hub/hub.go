package wsservicehub

import (
	wsserviceclient "first_socket/pkg/ws_service/client"
	wsservicerepository "first_socket/pkg/ws_service/repository"
	wsservicemessage "first_socket/pkg/ws_service/ws_message"

	"github.com/gorilla/websocket"
)

type hub struct {
	clientRepo wsservicerepository.IClientRepository
}

func (hub *hub) AddClient(
	ownerLogin string,
	connKey string,
	conn *websocket.Conn,
) (
	wsserviceclient.IWSClient,
	error,
) {
	client := hub.clientRepo.CreateClient(ownerLogin, connKey, conn)

	if err := hub.clientRepo.AddClient(client); err != nil {
		return nil, err
	}

	return client, nil
}

func (hub *hub) RemoveUser(ownerLogin string) {
	hub.clientRepo.RemoveUser(ownerLogin)
}

func (hub *hub) RemoveUserClient(ownerLogin string, connKey string) {
	hub.clientRepo.RemoveClient(ownerLogin, connKey)
}

func (hub *hub) SendUser(
	ownerLogin string,
	message wsservicemessage.IWSMessage,
) {
	clients := hub.clientRepo.GetUserClients(ownerLogin)

	hub.sendClients(clients, message)
}

func (hub *hub) SendUserWithoutClient(
	ownerLogin string,
	connKey string,
	message wsservicemessage.IWSMessage,
) {
	clients := hub.clientRepo.GetUserWithoutClient(ownerLogin, connKey)

	hub.sendClients(clients, message)
}

func (hub *hub) SendUsers(
	ownerLogins []string,
	message wsservicemessage.IWSMessage,
) {
	clients := hub.clientRepo.GetUsersClients(ownerLogins)

	hub.sendClients(clients, message)
}

func (hub *hub) sendClients(
	clients []wsserviceclient.IWSClient,
	message wsservicemessage.IWSMessage,
) {
	for _, client := range clients {
		client.Send(message)
	}
}
