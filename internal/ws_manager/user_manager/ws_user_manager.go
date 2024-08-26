package wsusermanager

import (
	"first_socket/internal/domain"
	"first_socket/internal/middleware"
	"first_socket/internal/repositories"
	wsmanager "first_socket/internal/ws_manager"

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

	manager.hub.AddClient(client)
	go manager.listen(client)
	client.Run()
}

func (manager *WSUserManager) listen(client *WSUserClient) {
	for {
		select {
		case message := <-client.receivedMessage:
			err := message.Map(
				func() error {
					innerClient, err := manager.hub.GetClientByName(message.Owner)
					if err != nil {
						return err
					}
					manager.hub.SendWithoutClient(
						innerClient,
						message,
					)
					return nil
				},
				func() error {
					innerClient, err := manager.hub.GetClientByName(message.Owner)
					if err != nil {
						return nil
					}

					manager.hub.RemoveClient(innerClient)
					manager.hub.SendAll(message)
					return nil
				},
			)

			if err != nil {
				break
			}
		}
	}

}

func NewWSUserManager(
	repository *repositories.UserRepository,
) *WSUserManager {
	return &WSUserManager{
		hub:        NewWSUserHub(),
		repository: repository,
	}
}
