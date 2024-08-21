package wsmanager

import (
	"first_socket/internal/domain"

	"github.com/gorilla/websocket"
)

type UserClient struct {
	User *domain.User
	Conn *websocket.Conn
	Hub  *UserHub
}

func (client *UserClient) Read() {
	defer func() {
		client.Hub.unregRequest <- client
		client.Conn.Close()
	}()

}

func newClient(
	user *domain.User,
	conn *websocket.Conn,
	hub *UserHub,
) *UserClient {
	return &UserClient{
		User: user,
		Conn: conn,
		Hub:  hub,
	}
}
