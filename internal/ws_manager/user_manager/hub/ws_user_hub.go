package wsuserhub

import (
	"first_socket/internal/repositories"
	wsclient "first_socket/internal/ws_manager/ws_client"
	wsmessage "first_socket/internal/ws_manager/ws_message"
)

type WSUserHub struct {
	clientRepository repositories.IClientRepository
}

func (hub *WSUserHub) AddClient(client wsclient.WSClient) {
	hub.clientRepository.AddClient(client)
}

func (hub *WSUserHub) RemoveClientByName(name string) {
	hub.clientRepository.RemoveClientByName(name)
}

func (hub *WSUserHub) GetClientsWithoutClientName(name string) []string {
	return hub.clientRepository.GetClientNamesWithoutClientName(name)
}

func (hub *WSUserHub) SendClientByName(
	name string,
	message wsmessage.WSMessage,
) {
	if client, err := hub.clientRepository.GetClientByName(name); err == nil {
		client.Send(message)
	}
}

func (hub *WSUserHub) SendWithoutClientName(
	name string,
	message wsmessage.WSMessage,
) {
	client := hub.clientRepository.GetClientsWithoutClientName(name)
	for _, c := range client {
		c.Send(message)
	}
}

func NewWSUserHub() *WSUserHub {
	return &WSUserHub{
		clientRepository: repositories.NewClientRepository(),
	}
}
