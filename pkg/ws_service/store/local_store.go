package wsservicestore

import (
	"errors"
	"first_socket/pkg/ws_service/client"

	"github.com/gorilla/websocket"
)

type localStore struct {
	clients map[string]map[string]wsserviceclient.IWSClient
}

func (store *localStore) AddClient(
	client wsserviceclient.IWSClient,
) error {
	login, connKey := client.GetOwnerLogin(), client.GetConnKey()

	if _, ok := store.clients[login][connKey]; ok {
		return errors.New("uncorrect connKey")
	}

	if _, ok := store.clients[login]; !ok {
		store.clients[login] = make(map[string]wsserviceclient.IWSClient)
	}

	store.clients[login][connKey] = client

	return nil
}

func (store *localStore) RemoveUser(login string) {
	if user, ok := store.clients[login]; ok {
		for _, client := range user {
			client.Close()
		}
		delete(user, login)
	}
}

func (store *localStore) RemoveClient(login string, connKey string) {
	if user, ok := store.clients[login]; ok {
		if client, ok := user[connKey]; ok {
			client.Close()
			delete(user, connKey)
		}

		if len(user) == 0 {
			delete(store.clients, login)
		}
	}
}

func (store *localStore) GetUserClients(login string) []wsserviceclient.IWSClient {
	clients := make([]wsserviceclient.IWSClient, 0)

	if user, ok := store.clients[login]; ok {
		for _, client := range user {
			clients = append(clients, client)
		}
	}

	return clients
}

func (store *localStore) GetUserWithoutClient(
	login string,
	connKey string,
) []wsserviceclient.IWSClient {
	clients := make([]wsserviceclient.IWSClient, 0)

	if user, ok := store.clients[login]; ok {
		for _, client := range user {
			if client.GetConnKey() != connKey {
				clients = append(clients, client)
			}
		}
	}

	return clients
}

func (store *localStore) GetUsersClients(logins []string) []wsserviceclient.IWSClient {
	clients := make([]wsserviceclient.IWSClient, 0)

	for _, login := range logins {
		if user, ok := store.clients[login]; ok {
			for _, client := range user {
				clients = append(clients, client)
			}
		}
	}

	return clients
}

func NewLocalStore() IStore {
	return &localStore{
		clients: make(map[string]map[string]wsserviceclient.IWSClient),
	}
}
