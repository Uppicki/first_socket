package wsmessage

import (
	"errors"
	wsresponses "first_socket/internal/ws_payload/responses"
)

type WSMessage struct {
	MessageType WSMessageType `json:"message_type"`
	Owner       string        `json:"owner_user"`
	Message     IMessage      `json:"message"`
}

func (message *WSMessage) MapToResponse() wsresponses.WSResponse {
	response := wsresponses.WSResponse{}

	response.ResponseType = wsresponses.WSResponseType(message.MessageType)

	switch message.MessageType {
	case Authorized:
		inner := wsresponses.WSSingleUserResponse{}
		inner.User = message.Owner
		response.Response = inner
	case Disauthorized:
		inner := wsresponses.WSSingleUserResponse{}
		inner.User = message.Owner
		response.Response = inner
	case Connected:
		inner := wsresponses.WSSingleUserResponse{}
		inner.User = message.Owner
		response.Response = inner
	case Disconnected:
		inner := wsresponses.WSSingleUserResponse{}
		inner.User = message.Owner
		response.Response = inner
	case UsersInfo:
		m, ok := message.Message.(UsersInfoMessage)
		if !ok {
			break
		}
		inner := wsresponses.WSUsersResponse{}
		inner.Users = m.Users
		response.Response = inner
	case ChatsInfo:
		m, ok := message.Message.(ChatsInfoMessage)
		if !ok {
			break
		}
		inner := wsresponses.WSChatsResponse{}
		inner.Chats = m.Chats
		response.Response = inner
	case ChatMessages:
		m, ok := message.Message.(ChatMessagesMessage)
		if !ok {
			break
		}
		inner := wsresponses.WSChatMessagesResponse{}
		inner.Companion = m.Companion
		inner.Chat = m.Chat
		inner.Messages = m.Messages
		response.Response = inner
	case MessageSend:
		m, ok := message.Message.(MessageSendMessage)
		if !ok {
			break
		}
		inner := wsresponses.MessageSendResponse{}
		inner.IsSended = m.IsSended
		inner.Message = m.Message
		response.Response = inner
	case MessageNotification:
		m, ok := message.Message.(MessageNotificationMessage)
		if !ok {
			break
		}
		inner := wsresponses.MessageNotificationResponse{}
		inner.Owner = message.Owner
		inner.Message = m.Message
		response.Response = inner
	}

	return response
}

func (message *WSMessage) MapHandler(
	connectedFunc func(WSMessage),
	disconnectedFunc func(WSMessage),
	usersInfoFunc func(WSMessage, UsersInfoMessage),
	chatsInfoFunc func(WSMessage, ChatsInfoMessage),
	chatMessagesFunc func(WSMessage, ChatMessagesMessage),
	messageSendFunc func(WSMessage, MessageSendMessage),
) error {
	var err error

	switch message.MessageType {
	case Connected:
		connectedFunc(*message)
	case Disconnected:
		disconnectedFunc(*message)
	case UsersInfo:
		if innerMessage, ok := message.Message.(UsersInfoMessage); ok {
			usersInfoFunc(*message, innerMessage)
		} else {
			return errors.New("dawncasting error")
		}
	case ChatsInfo:
		if innerMessage, ok := message.Message.(ChatsInfoMessage); ok {
			chatsInfoFunc(*message, innerMessage)
		} else {
			return errors.New("dawncasting error")
		}
	case ChatMessages:
		if innerMessage, ok := message.Message.(ChatMessagesMessage); ok {
			chatMessagesFunc(*message, innerMessage)
		} else {
			return errors.New("dawncasting error")
		}
	case MessageSend:
		if innerMessage, ok := message.Message.(MessageSendMessage); ok {
			messageSendFunc(*message, innerMessage)
		} else {
			return errors.New("dawncasting error")
		}
	default:
		err = errors.New("unsresolved wsmessage type")
	}

	return err
}
