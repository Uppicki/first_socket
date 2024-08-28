package wsrequests

type WSRequestType string

const (
	Disconnected WSRequestType = "DisconnectedType"
	ChatsInfo    WSRequestType = "ChatsInfoType"
	UsersInfo    WSRequestType = "UsersInfoType"
	ChatMessages WSRequestType = "ChatMessageType"
	MessageSend  WSRequestType = "MessageSendType"
)
