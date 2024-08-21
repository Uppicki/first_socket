package wsmanager

type UserHub struct {
	clients      map[string]*UserClient
	regRequest   chan *UserClient
	unregRequest chan *UserClient
}

func (hub *UserHub) Run() {
	go hub.run()
}

func (hub *UserHub) regNewClient(client *UserClient) {
	hub.clients[client.User.Name] = client
}

func (hub *UserHub) unregNewClient(client *UserClient) {
	if _, ok := hub.clients[client.User.Name]; ok {
		delete(hub.clients, client.User.Name)
	}
}

func (hub *UserHub) run() {
	for {
		select {
		case req := <-hub.regRequest:
			hub.regNewClient(req)
		case req := <-hub.regRequest:
			hub.unregNewClient(req)
		}
	}
}

func NewHub() *UserHub {
	return &UserHub{
		clients:      make(map[string]*UserClient),
		regRequest:   make(chan *UserClient),
		unregRequest: make(chan *UserClient),
	}
}
