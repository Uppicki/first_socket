package wsmessage

type IMessage interface {
}

type UsersInfoMessage struct {
	AuthorizedUSers []string
	ConnectedUsers  map[string]bool
}
