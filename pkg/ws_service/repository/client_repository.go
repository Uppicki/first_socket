package wsservicerepository

import (
	client "first_socket/pkg/ws_service/client"
	wsservicestore "first_socket/pkg/ws_service/store"

	"github.com/gorilla/websocket"
)

type clientRepository struct {
	store wsservicestore.IStore
}

func (repo *clientRepository) CreateClient(
	ownerLogin string,
	connKey string,
	conn *websocket.Conn,
) client.IWSClient {
	return client.NewWSClient(ownerLogin, connKey, conn)
}

func (repo *clientRepository) AddClient(
	client client.IWSClient,
) error {
	return repo.AddClient(client)
}

func (repo *clientRepository) RemoveUser(
	ownerLogin string,
) {
	repo.store.RemoveUser(ownerLogin)
}

func (repo *clientRepository) RemoveClient(
	ownerLogin string,
	connKey string,
) {
	repo.store.RemoveClient(ownerLogin, connKey)
}

func (repo *clientRepository) GetUserClients(
	ownerLogin string,
) []client.IWSClient {
	return repo.store.GetUserClients(ownerLogin)
}

func (repo *clientRepository) GetUserWithoutClient(
	ownerLogin string,
	connKey string,
) []client.IWSClient {
	return repo.store.GetUserWithoutClient(ownerLogin, connKey)
}

func (repo *clientRepository) GetUsersClients(
	logins []string,
) []client.IWSClient {
	return repo.GetUsersClients(logins)
}
