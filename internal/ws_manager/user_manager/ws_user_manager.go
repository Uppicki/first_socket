package wsusermanager

import (
	"first_socket/internal/domain"
	"first_socket/internal/middleware"
	"first_socket/internal/repositories"
	wsmanager "first_socket/internal/ws_manager"
	wsmessage "first_socket/internal/ws_manager/ws_message"

	"github.com/gin-gonic/gin"
)

type WSUserManager struct {
	hub            *WSUserHub
	repository     *repositories.UserRepository
	chatRepository *repositories.ChatRepository
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

					selfMessage := manager.createUsersInfoMessage(
						innerClient,
					)

					manager.hub.SendClient(
						innerClient,
						selfMessage,
					)

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
				func() error {
					return nil
				},
			)

			if err != nil {
				break
			}
		}
	}

}

func (manager *WSUserManager) createUsersInfoMessage(
	client *WSUserClient,
) wsmessage.WSMessage {
	users := manager.repository.GetUsernamesWithoutUser(
		client.user.Name,
	)
	connectedUsers := manager.hub.GetClientNamesWithoutClient(
		client,
	)

	message := wsmessage.UsersInfoMessage{
		AuthorizedUSers: users,
		ConnectedUsers:  connectedUsers,
	}

	return wsmessage.WSMessage{
		MessageType: wsmessage.UsersInfoType,
		Owner:       client.user.Name,
		Message:     message,
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
