package wsmessage

type WSMessage struct {
	MessageType WSMessageType
	Owner       string
	TargetUser  string
	Message     string
}
