package wsusermanager

type WSUserHub struct {
	clients map[string]*WSUserClient
}

func (hub *WSUserHub) addClient(client *WSUserClient) {
	hub.clients[client.GetNameOwnerUser()] = client
}

func NewWSUserHub() *WSUserHub {
	return &WSUserHub{
		clients: make(map[string]*WSUserClient),
	}
}
