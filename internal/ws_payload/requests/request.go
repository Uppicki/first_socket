package wsrequests

import (
	"encoding/json"
	"errors"
	wsmessage "first_socket/internal/ws_manager/ws_message"
)

type WSRequest struct {
	Type    WSRequestType   `json:"request_type"`
	Request json.RawMessage `json:"request"`
}

func (req *WSRequest) MapToMessage(
	disconnectedFunc func(WSRequest) (*wsmessage.WSMessage, error),
	chatsInfoFunc func(WSRequest) (*wsmessage.WSMessage, error),
	usersInfoFunc func(WSRequest) (*wsmessage.WSMessage, error),
	chatMessagesFunc func(WSRequest, WSChatMessagesRequest) (*wsmessage.WSMessage, error),
	messageSendFunc func(WSRequest, WSSendMessageRequest) (*wsmessage.WSMessage, error),
) (*wsmessage.WSMessage, error) {
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
			return nil, err
		}
		return chatMessagesFunc(*req, r)
	case MessageSend:
		var r WSSendMessageRequest
		if err := json.Unmarshal(req.Request, &r); err != nil {
			return nil, err
		}
		return messageSendFunc(*req, r)
	default:
		return nil, errors.New("Undefined wsrequest type")
	}
}
