package wsservice

import (
	wsserviceclient "first_socket/pkg/ws_service/client"
	"first_socket/pkg/ws_service/hub"
	"net/http"

	"github.com/gorilla/websocket"
)

type WSService struct {
	hub wsservicehub.IWSClientHub
}

func (service *WSService) ServeWS(
	owner string,
	connKey string,
	payload Payload,
) error {
	conn, err := service.CreateConnection(
		payload.Writer,
		payload.Request,
		payload.Header,
	)
	if err != nil {
		return err
	}

	client, clientErr := service.hub.AddClient(owner, connKey, conn)
	if clientErr != nil {
		return err
	}

	go service.Listen(client)
	client.Run()

	return nil
}

func (service *WSService) Listen(client wsserviceclient.IWSClient) {
	channel := client.GetReceivedChan()
	for {
		select {
		case message := <-channel:
			message.Map()
		}
	}
}

func (service *WSService) CreateConnection(
	writer http.ResponseWriter,
	request *http.Request,
	header http.Header,
) (*websocket.Conn, error) {
	return Upgrader.Upgrade(writer, request, header)
}
