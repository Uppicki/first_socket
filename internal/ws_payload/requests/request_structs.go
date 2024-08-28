package wsrequests

type WSChatMessagesRequest struct {
	Companion string `json:"companion"`
}

type WSSendMessageRequest struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"message_text"`
}
