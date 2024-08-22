package wsusermanager

import (
	"first_socket/internal/repositories"
	wsmanager "first_socket/internal/ws_manager"

	"github.com/gin-gonic/gin"
)

type WSUserManager struct {
	hub        *WSUserHub
	repository *repositories.UserRepository
}

func (manager *WSUserManager) ServeWS(ctx *gin.Context) {

	user, err := manager.repository.GetUserByName("vovav")
	if err != nil {
		return
	}

	conn, err := wsmanager.Upgrader.Upgrade(
		ctx.Writer,
		ctx.Request,
		nil,
	)
	if err != nil {
		return
	}

	client := NewWSUserClient(conn, user)

	manager.hub.addClient(client)
}

func NewWSUserManager(
	repository *repositories.UserRepository,
) *WSUserManager {
	return &WSUserManager{
		hub:        NewWSUserHub(),
		repository: repository,
	}
}
