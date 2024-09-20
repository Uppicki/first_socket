package repositories

import (
	"errors"
	wsclient "first_socket/internal/ws_manager/ws_client"
)

type clientRepository struct {
	clients map[string]wsclient.WSClient
}

func (repo *clientRepository) AddClient(client wsclient.WSClient) {
	repo.clients[client.GetNameOwnerUser()] = client
}

func (repo *clientRepository) RemoveClientByName(name string) {
	if client, ok := repo.clients[name]; ok {
		delete(repo.clients, name)
		client.Close()
	}
}

func (repo *clientRepository) GetClientNamesWithoutClientName(name string) []string {
	res := make([]string, len(repo.clients)-1)

	for _, client := range repo.clients {
		res = append(res, client.GetNameOwnerUser())
	}

	return res
}

func (repo *clientRepository) GetClientByName(name string) (
	wsclient.WSClient,
	error,
) {
	if client, ok := repo.clients[name]; ok {
		return client, nil
	} else {
		return nil, errors.New("client doesn`t exsists")
	}
}

func (repo *clientRepository) GetClientsWithoutClientName(name string) []wsclient.WSClient {
	res := make([]wsclient.WSClient, len(repo.clients)-1)

	for _, client := range repo.clients {
		res = append(res, client)
	}

	return res
}

func NewClientRepository() *clientRepository {
	return &clientRepository{
		clients: make(map[string]wsclient.WSClient),
	}
}
