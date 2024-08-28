package wsuserclient

import (
	"first_socket/internal/domain"
	wsmessage "first_socket/internal/ws_manager/ws_message"
	wsrequests "first_socket/internal/ws_payload/requests"
	"sync"

	"github.com/gorilla/websocket"
)

type WSUserClient struct {
	conn            *websocket.Conn
	user            *domain.User
	sendMessage     chan wsmessage.WSMessage
	receivedMessage chan wsmessage.WSMessage
	isActive        bool
	mu              sync.Mutex
}

func (client *WSUserClient) GetNameOwnerUser() string {
	return client.user.Name
}

func (client *WSUserClient) Run() {
	client.mu.Lock()
	client.isActive = true
	client.mu.Unlock()
	go client.readerRun()
	go client.writerRun()
}

func (client *WSUserClient) Close() {
	client.mu.Lock()
	defer client.mu.Unlock()

	client.isActive = false
	client.conn.Close()
	close(client.sendMessage)
	close(client.receivedMessage)
}

func (client *WSUserClient) Send(message wsmessage.WSMessage) {
	select {
	case client.sendMessage <- message:
	default:
	}
}

func (client *WSUserClient) GetReceivedChan() <-chan wsmessage.WSMessage {
	return client.receivedMessage
}

func (client *WSUserClient) readerRun() {
	defer func() {
		client.mu.Lock()
		if client.isActive {
			client.receivedMessage <- wsmessage.WSMessage{
				MessageType: wsmessage.Disconnected,
				Owner:       client.user.Name,
			}
			client.isActive = false
		}
		client.mu.Unlock()
	}()

	client.receivedMessage <- wsmessage.WSMessage{
		MessageType: wsmessage.Connected,
		Owner:       client.user.Name,
	}

	for {
		client.mu.Lock()
		if !client.isActive {
			client.mu.Unlock()
			break
		}
		client.mu.Unlock()

		var req wsrequests.WSRequest

		err := client.conn.ReadJSON(&req)

		if err != nil {
			break
		}

		if message, innerErr := req.MapToMessage(
			client.disconnectedFunc,
			client.chatsInfoFunc,
			client.usersInfoFunc,
			client.chatMessagesFunc,
			client.messageSendFunc,
		); innerErr == nil {
			client.receivedMessage <- *message
		}

	}
}

func (client *WSUserClient) disconnectedFunc(
	request wsrequests.WSRequest,
) (
	*wsmessage.WSMessage,
	error,
) {
	message := &wsmessage.WSMessage{
		MessageType: wsmessage.WSMessageType(request.Type),
		Owner:       client.user.Name,
	}

	return message, nil
}

func (client *WSUserClient) chatsInfoFunc(
	request wsrequests.WSRequest,
) (
	*wsmessage.WSMessage,
	error,
) {
	message := &wsmessage.WSMessage{
		MessageType: wsmessage.WSMessageType(request.Type),
		Owner:       client.user.Name,
		Message:     wsmessage.ChatsInfoMessage{},
	}

	return message, nil
}

func (client *WSUserClient) usersInfoFunc(
	request wsrequests.WSRequest,
) (
	*wsmessage.WSMessage,
	error,
) {
	message := &wsmessage.WSMessage{
		MessageType: wsmessage.WSMessageType(request.Type),
		Owner:       client.user.Name,
		Message:     wsmessage.UsersInfoMessage{},
	}

	return message, nil
}

func (client *WSUserClient) chatMessagesFunc(
	request wsrequests.WSRequest,
	chatMessageRequest wsrequests.WSChatMessagesRequest,
) (
	*wsmessage.WSMessage,
	error,
) {
	message := &wsmessage.WSMessage{
		MessageType: wsmessage.WSMessageType(request.Type),
		Owner:       client.user.Name,
		Message: wsmessage.ChatMessagesMessage{
			Companion: chatMessageRequest.Companion,
		},
	}

	return message, nil
}

func (client *WSUserClient) messageSendFunc(
	request wsrequests.WSRequest,
	sendMessageRequest wsrequests.WSSendMessageRequest,
) (
	*wsmessage.WSMessage,
	error,
) {
	message := &wsmessage.WSMessage{
		MessageType: wsmessage.WSMessageType(request.Type),
		Owner:       client.user.Name,
		Message: wsmessage.MessageSendMessage{
			IsSended: false,
			Message:  sendMessageRequest.Text,
			ChatID:   sendMessageRequest.ChatId,
		},
	}

	return message, nil
}

func (client *WSUserClient) writerRun() {
	for {
		client.mu.Lock()
		if !client.isActive {
			client.mu.Unlock()
			return
		}
		client.mu.Unlock()

		if message, ok := <-client.sendMessage; ok {
			response := message.MapToResponse()

			client.conn.WriteJSON(response)
		}
	}
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
