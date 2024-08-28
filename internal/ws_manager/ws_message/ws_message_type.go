package wsmessage

type WSMessageType string

const (
	Authorized          WSMessageType = "AuthorizedType"    //only send
	Disauthorized       WSMessageType = "DisauthorizedType" // only send
	Connected           WSMessageType = "ConnectedType"     // only send
	Disconnected        WSMessageType = "DisconnectedType"
	UsersInfo           WSMessageType = "UsersInfoType"
	ChatsInfo           WSMessageType = "ChatsInfoType"
	ChatMessages        WSMessageType = "ChatMessagesType"
	MessageSend         WSMessageType = "MessangeSendType"
	MessageNotification WSMessageType = "MessageNotificationType" //only send
)
