package wsmessage

type IMessage interface {
}

type UsersInfoMessage struct {
	Users []string
}

type ChatsInfoMessage struct {
	Chats []string
}

type ChatMessagesMessage struct {
	Chat      string
	Companion string
	Messages  []string
}

type MessageSendMessage struct {
	ChatID   string
	IsSended bool
	Message  string
}

type MessageNotificationMessage struct {
	Message string
}
