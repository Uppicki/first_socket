package wsrequests

import (
	"encoding/json"
	"errors"
)

type WSRequest struct {
	Type    WSRequestType   `json:"request_type"`
	Request json.RawMessage `json:"request"`
}

func (req *WSRequest) Map(
	disconnectedFunc func(WSRequest) error,
	chatsInfoFunc func(WSRequest) error,
	usersInfoFunc func(WSRequest) error,
	chatMessagesFunc func(WSRequest, WSChatMessagesRequest) error,
	messageSendFunc func(WSRequest, WSSendMessageRequest) error,
) error {
	switch req.Type {
	case Disconnected:
		return disconnectedFunc(*req)
	case ChatsInfo:
		return chatsInfoFunc(*req)
	case UsersInfo:
		return usersInfoFunc(*req)
	case ChatMessages:
		var r WSChatMessagesRequest
		if err := json.Unmarshal(req.Request, &r); err != nil {
			return err
		}
		return chatMessagesFunc(*req, r)
	case MessageSend:
		var r WSSendMessageRequest
		if err := json.Unmarshal(req.Request, &r); err != nil {
			return err
		}
		return messageSendFunc(*req, r)
	default:
		return errors.New("Undefined wsrequest type")
	}
}
