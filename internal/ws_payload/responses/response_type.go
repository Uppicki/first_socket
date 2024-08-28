package wsresponses

type WSResponseType string

const (
	Authorized          WSResponseType = "AuthorizedType"
	Disauthorized       WSResponseType = "DisauthorizedType"
	Connected           WSResponseType = "ConnectedType"
	Disconnected        WSResponseType = "DisconnectedType"
	UsersInfo           WSResponseType = "UsersInfoType"
	ChatsInfo           WSResponseType = "ChatsInfoType"
	ChatMessages        WSResponseType = "ChatMessagesType"
	MessageSend         WSResponseType = "MessangeSendType"
	MessageNotification WSResponseType = "MessageNotificationType"
)
