package wsusermanager

import (
	"first_socket/internal/domain"
	wsmessage "first_socket/internal/ws_manager/ws_message"
	wsresponses "first_socket/internal/ws_payload/responses"
	"fmt"

	"github.com/gorilla/websocket"
)

type WSUserClient struct {
	conn           *websocket.Conn
	user           *domain.User
	sendingMessage chan *wsmessage.WSMessage
}

func (client *WSUserClient) GetNameOwnerUser() string {
	return client.user.Name
}

func (client *WSUserClient) readerRun() {
	defer func() {
		client.conn.Close()
	}()

	for {

	}
}

func (client *WSUserClient) writerRun() {
	defer func() {
		client.conn.Close()
	}()

	fmt.Println("asdasdasd")

	for {
		message, ok := <-client.sendingMessage
		fmt.Println(ok)
		if ok {
			if err := client.conn.WriteJSON(
				wsresponses.ConnectedResponse{
					Username:      message.Owner,
					ConnectedType: wsresponses.ConnectedResponseType(message.Message),
				},
			); err != nil {
				fmt.Println(err)
			}
		}
	}

}

func (client *WSUserClient) run() {
	go client.readerRun()
	go client.writerRun()
}

func NewWSUserClient(
	conn *websocket.Conn,
	user *domain.User,
) *WSUserClient {
	return &WSUserClient{
		conn:           conn,
		user:           user,
		sendingMessage: make(chan *wsmessage.WSMessage),
	}
}
