package wsusermanager

import (
	"first_socket/internal/domain"

	"github.com/gorilla/websocket"
)

type WSUserClient struct {
	conn *websocket.Conn
	user *domain.User
}

func (client *WSUserClient) GetNameOwnerUser() string {
	return client.user.Name
}

func NewWSUserClient(
	conn *websocket.Conn,
	user *domain.User,
) *WSUserClient {
	return &WSUserClient{
		conn: conn,
		user: user,
	}
}
