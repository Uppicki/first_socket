package wsusermanager

import (
	"errors"
	wsmessage "first_socket/internal/ws_manager/ws_message"

)

type WSUserHub struct {
	clients map[string]*WSUserClient
}

func (hub *WSUserHub) AddClient(client *WSUserClient) {
	hub.clients[client.GetNameOwnerUser()] = client
}

func (hub *WSUserHub) GetClientByName(name string) (*WSUserClient, error) {
	client, ok := hub.clients[name]
	if !ok {
		return nil, errors.New("Client doesn`t exsist")
	}
	return client, nil
}

func (hub *WSUserHub) RemoveClient(client *WSUserClient) {
	if _, ok := hub.clients[client.user.Name]; ok {
		delete(hub.clients, client.user.Name)
		client.Close()
	}
}

func (hub *WSUserHub) RemoveClientByName(name string) {
	client, err := hub.GetClientByName(name)

	if err == nil {
		hub.RemoveClient(client)
	}
}

func (hub *WSUserHub) SendAll(
	message wsmessage.WSMessage,
) {
	for _, cl := range hub.clients {
		cl.sendMessage <- message
	}
}

func (hub *WSUserHub) SendWithoutClient(
	client *WSUserClient,
	message wsmessage.WSMessage,
) {
	for _, cl := range hub.clients {
		if cl.user.Name != client.user.Name {
			cl.sendMessage <- message
		}
	}
}

func NewWSUserHub() *WSUserHub {
	return &WSUserHub{
		clients: make(map[string]*WSUserClient),
	}
}
