package repositories

import wsclient "first_socket/internal/ws_manager/ws_client"

type IClientRepository interface {
	AddClient(wsclient.WSClient)
	GetClientByName(string) wsclient.WSClient
	GetClients() []wsclient.WSClient
}

type IChatRepository interface {
	RemoveUserChats(string)
	GetUserChats(string) []string
	GetChat(string, string) string
	GetChatMessages(string) []string
	WriteMessage(string, string)
}
