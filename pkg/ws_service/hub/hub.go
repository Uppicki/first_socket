package wsservicehub

import (
	"first_socket/pkg/ws_service/client"
	"first_socket/pkg/ws_service/repository"

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
	hub.clientRepo.AddClient(ownerLogin, connKey, conn)

	return nil, nil
}
