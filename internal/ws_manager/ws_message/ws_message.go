package wsmessage

import "errors"

type WSMessage struct {
	MessageType WSMessageType `json:"message_type"`
	Owner       string        `json:"owner_user"`
	TargetUser  string        `json:"target_user"`
	Message     string        `json:"message"`
}

func (message *WSMessage) Map(
	connectedFunc func() error,
	disconnectedFunc func() error,
) error {
	switch message.MessageType {
	case ConnectedType:
		return connectedFunc()
	case DisconnectedType:
		return disconnectedFunc()
	default:
		return errors.New("Undefined MessageType")
	}
}
