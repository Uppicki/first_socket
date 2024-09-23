package repositories

import (
	"first_socket/internal/domain"
	wsclient "first_socket/internal/ws_manager/ws_client"
)

type IUserRepository interface {
	IsLoginExsist(string) bool
	CreateUser(user domain.User) error
}

type IClientRepository interface {
	AddClient(wsclient.WSClient)
	RemoveClientByName(string)
	GetClientNamesWithoutClientName(string) []string
	GetClientByName(string) (wsclient.WSClient, error)
	GetClientsWithoutClientName(string) []wsclient.WSClient
}

type IChatRepository interface {
	RemoveUserChats(string)
	GetUserChats(string) []string
	GetChat(string, string) string
	GetChatMessages(string) []string
	WriteMessage(string, string)
}
