package wsmessage

type WSMessageType string

const (
	AuthorizedType    WSMessageType = "AuthorizedType"
	DisauthorizedType WSMessageType = "DisauthorizedType"
	ConnectedType     WSMessageType = "ConnectedType"
	DisconnectedType  WSMessageType = "DisconnectedType"
	UsersInfoType     WSMessageType = "UsersInfoType"
	UserMessageType   WSMessageType = "UserMessageType"
)
