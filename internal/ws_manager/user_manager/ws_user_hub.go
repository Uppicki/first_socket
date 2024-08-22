package wsusermanager

import (
	wsmessage "first_socket/internal/ws_manager/ws_message"
	"fmt"
)

type WSUserHub struct {
	clients map[string]*WSUserClient
}

func (hub *WSUserHub) addClient(client *WSUserClient) {
	hub.clients[client.GetNameOwnerUser()] = client
	client.run()
	hub.notifyWithoutClient(client)
}

func (hub *WSUserHub) notifyWithoutClient(client *WSUserClient) {
	for _, cl := range hub.clients {
		if cl.user.Name != client.user.Name {
			fmt.Println("asdasd")
			cl.sendingMessage <- &wsmessage.WSMessage{
				Owner: client.GetNameOwnerUser(),
			}
		}
	}
}

func NewWSUserHub() *WSUserHub {
	return &WSUserHub{
		clients: make(map[string]*WSUserClient),
	}
}
