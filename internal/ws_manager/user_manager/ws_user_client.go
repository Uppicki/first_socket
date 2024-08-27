package wsusermanager

import (
	"first_socket/internal/domain"
	wsmessage "first_socket/internal/ws_manager/ws_message"
	wsresponses "first_socket/internal/ws_payload/responses"
	"fmt"

	"github.com/gorilla/websocket"
)

type WSUserClient struct {
	conn            *websocket.Conn
	user            *domain.User
	sendMessage     chan wsmessage.WSMessage
	receivedMessage chan wsmessage.WSMessage
	isActive        bool
}

func (client *WSUserClient) GetNameOwnerUser() string {
	return client.user.Name
}

func (client *WSUserClient) readerRun() {
	defer func() {
		if client.isActive {
			client.receivedMessage <- wsmessage.WSMessage{
				MessageType: wsmessage.DisconnectedType,
				Owner:       client.user.Name,
			}
			client.isActive = false
		}
	}()

	client.receivedMessage <- wsmessage.WSMessage{
		MessageType: wsmessage.ConnectedType,
		Owner:       client.user.Name,
	}

	for {
		var message wsmessage.WSMessage

		err := client.conn.ReadJSON(&message)

		if err != nil {
			fmt.Println("asd")
			break
		}

		fmt.Println(message)

		message.Owner = client.user.Name

		client.receivedMessage <- message

	}
}

func (client *WSUserClient) writerRun() {
	defer func() {

	}()

	for {
		message, ok := <-client.sendMessage
		if ok {
			message.Map(
				// Connected response
				func() error {
					innerErr := client.conn.WriteJSON(
						wsresponses.ConnectedResponse{
							Username:      message.Owner,
							ConnectedType: wsresponses.ConnectedResponseType(message.MessageType),
						},
					)
					return innerErr
				},
				// Disconnected response
				func() error {
					innerErr := client.conn.WriteJSON(
						wsresponses.ConnectedResponse{
							Username:      message.Owner,
							ConnectedType: wsresponses.ConnectedResponseType(message.MessageType),
						},
					)
					return innerErr
				},
				// Users info response
				func() error {
					m, innerOk := message.Message.(wsmessage.UsersInfoMessage)

					if !innerOk {
						return nil
					}

					innerErr := client.conn.WriteJSON(
						wsresponses.NewUsersInfoResponse(
							m.AuthorizedUSers,
							m.ConnectedUsers,
						),
					)

					return innerErr
				},
			)
		}

	}
}

func (client *WSUserClient) Run() {
	client.isActive = true
	go client.readerRun()
	go client.writerRun()
}

func (client *WSUserClient) Close() {
	client.isActive = false
	client.conn.Close()
	close(client.sendMessage)

}

func NewWSUserClient(
	conn *websocket.Conn,
	user *domain.User,
) *WSUserClient {
	return &WSUserClient{
		conn:            conn,
		user:            user,
		sendMessage:     make(chan wsmessage.WSMessage),
		receivedMessage: make(chan wsmessage.WSMessage),
		isActive:        false,
	}
}
