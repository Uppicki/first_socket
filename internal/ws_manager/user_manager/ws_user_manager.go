package wsusermanager

import (
	"first_socket/internal/domain"
	"first_socket/internal/middleware"
	"first_socket/internal/repositories"
	wsmanager "first_socket/internal/ws_manager"
	wsuserclient "first_socket/internal/ws_manager/user_manager/client"
	wsuserhub "first_socket/internal/ws_manager/user_manager/hub"
	wshub "first_socket/internal/ws_manager/ws_hub"
	wsmessage "first_socket/internal/ws_manager/ws_message"

	"github.com/gin-gonic/gin"
)

type WSUserManager struct {
	hub            wshub.WSHub
	repository     *repositories.UserRepository
	chatRepository repositories.IChatRepository
}

func (manager *WSUserManager) NotifyAboutAuthorized(username string) {
	message := wsmessage.WSMessage{
		MessageType: wsmessage.Disauthorized,
		Owner:       username,
	}

	manager.hub.SendWithoutClientName(username, message)
}

func (manager *WSUserManager) Disauthorize(username string) {
	manager.chatRepository.RemoveUserChats(username)

	message := wsmessage.WSMessage{
		MessageType: wsmessage.Disauthorized,
		Owner:       username,
	}

	manager.hub.RemoveClientByName(username)
	manager.hub.SendWithoutClientName(username, message)
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

	client := wsuserclient.NewWSUserClient(conn, user)

	manager.hub.AddClient(client)
	go manager.listen(client)
	client.Run()
}

func (manager *WSUserManager) listen(client *wsuserclient.WSUserClient) {
	channel := client.GetReceivedChan()
	for {
		select {
		case message := <-channel:

			message.MapHandler(
				manager.connectedHandler,
				manager.disconnectedHandler,
				manager.usersInfoHandler,
				manager.chatsInfoHandler,
				manager.chatMessagesHandler,
				manager.messageSendHandler,
			)
		}
	}

}

func (manager *WSUserManager) connectedHandler(
	message wsmessage.WSMessage,
) {
	manager.hub.SendWithoutClientName(message.Owner, message)
}

func (manager *WSUserManager) disconnectedHandler(
	message wsmessage.WSMessage,
) {
	manager.hub.RemoveClientByName(message.Owner)
	manager.hub.SendWithoutClientName(message.Owner, message)
}

func (manager *WSUserManager) usersInfoHandler(
	message wsmessage.WSMessage,
	innerMessage wsmessage.UsersInfoMessage,
) {
	innerMessage.Users = manager.hub.GetClientsWithoutClientName(message.Owner)
	message.Message = innerMessage
	manager.hub.SendClientByName(message.Owner, message)
}

func (manager *WSUserManager) chatsInfoHandler(
	message wsmessage.WSMessage,
	innerMessage wsmessage.ChatsInfoMessage,
) {
	innerMessage.Chats = manager.chatRepository.GetUserChats(message.Owner)
	message.Message = innerMessage
	manager.hub.SendClientByName(message.Owner, message)
}

func (manager *WSUserManager) chatMessagesHandler(
	message wsmessage.WSMessage,
	innerMessage wsmessage.ChatMessagesMessage,
) {
	innerMessage.Chat = manager.chatRepository.GetChat(
		message.Owner,
		innerMessage.Companion,
	)
	innerMessage.Messages = manager.chatRepository.GetChatMessages(
		innerMessage.Chat,
	)
	message.Message = innerMessage
	manager.hub.SendClientByName(message.Owner, message)
}

func (manager *WSUserManager) messageSendHandler(
	message wsmessage.WSMessage,
	innerMessage wsmessage.MessageSendMessage,
) {
	manager.chatRepository.WriteMessage(innerMessage.ChatID, innerMessage.Message)
	innerMessage.IsSended = true
	message.Message = innerMessage
	manager.hub.SendClientByName(message.Owner, message)

	im := wsmessage.MessageNotificationMessage{
		Message: innerMessage.Message,
	}
	m := wsmessage.WSMessage{
		MessageType: wsmessage.MessageNotification,
		Owner:       message.Owner,
		Message:     im,
	}
	manager.hub.SendClientByName(innerMessage.Companion, m)
}

func NewWSUserManager(
	userRepo *repositories.UserRepository,
	chatRepo *repositories.IChatRepository,
) *WSUserManager {
	return &WSUserManager{
		hub:            wsuserhub.NewWSUserHub(),
		repository:     userRepo,
		chatRepository: *chatRepo,
	}
}
