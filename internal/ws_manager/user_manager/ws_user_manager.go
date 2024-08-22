package wsusermanager

import (
	"first_socket/internal/domain"
	"first_socket/internal/middleware"
	"first_socket/internal/repositories"
	wsmanager "first_socket/internal/ws_manager"
	"fmt"

	"github.com/gin-gonic/gin"
)

type WSUserManager struct {
	hub        *WSUserHub
	repository *repositories.UserRepository
}

func (manager *WSUserManager) ServeWS(ctx *gin.Context) {

	ctxUser, _ := ctx.Get(middleware.UserContextKey)

	user, ok := ctxUser.(*domain.User)

	if !ok {
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
	fmt.Println("asd")
}

func NewWSUserManager(
	repository *repositories.UserRepository,
) *WSUserManager {
	return &WSUserManager{
		hub:        NewWSUserHub(),
		repository: repository,
	}
}
