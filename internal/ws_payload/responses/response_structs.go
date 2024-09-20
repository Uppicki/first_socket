package wsresponses

type IWSResponse interface {
}

type WSSingleUserResponse struct {
	User string `json:"username"`
}

type WSUsersResponse struct {
	Users []string `json:"users"`
}

type WSChatsResponse struct {
	Chats []string `json:"chats"`
}

type WSChatMessagesResponse struct {
	Companion string   `json:"companion"`
	Chat      string   `json:"chat"`
	Messages  []string `json:"messages"`
}

type MessageSendResponse struct {
	IsSended bool   `json:"is_sended"`
	Message  string `json:"message"`
}

type MessageNotificationResponse struct {
	Owner   string `json:"owner"`
	Message string `json:"message"`
}

type ChatDeleteNotificationResponse struct {
	Chat string `json:"chat_id"`
}
