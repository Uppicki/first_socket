package wsrequests

type WSSendMessageRequest struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"message_text"`
}
