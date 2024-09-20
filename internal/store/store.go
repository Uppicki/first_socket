package store

import (
	"errors"
	"first_socket/internal/domain"
)

type Store struct {
	users          map[string]*domain.User
	chats          map[string]*domain.Chat
	realationships map[string]map[string]string
}

func (store *Store) AddUser(name string) {
	store.users[name] = &domain.User{Name: name}
}

func (store *Store) GetAllUsers() map[string]*domain.User {
	res := make(map[string]*domain.User, len(store.users))

	for _, user := range store.users {
		res[user.Name] = user
	}

	return res
}

func (store *Store) GetUserByName(
	name string,
) (*domain.User, error) {
	user, ok := store.users[name]
	if !ok {
		return nil, errors.New("User doesn`t authorized")
	}
	return user, nil
}

func (store *Store) GetChatByUserNames(username1, username2 string) (*domain.Chat, error) {
	user1, ok1 := store.GetUserByName(username1)
	user2, ok2 := store.GetUserByName(username2)

	if (ok1 != nil) || (ok2 != nil) {
		return nil, errors.New("User doesn`t authorized")
	}

	chat, err := store.getChat(user1, user2)
	if err != nil {
		chat = domain.NewChat(user1, user2)
		store.realationships[user1.Name][user2.Name] = chat.ID
		store.realationships[user2.Name][user1.Name] = chat.ID
	}

	return chat, nil
}

func (store *Store) getChat(user1, user2 *domain.User) (*domain.Chat, error) {
	chat, ok := store.realationships[user1.Name][user2.Name]
	if !ok {
		return nil, errors.New("Chat doesn`t exsist")
	}

	return store.chats[chat], nil
}

func (store *Store) RemoveUserChats(name string) {
	users := make([]string, 0)
	chats := make([]string, 0)

	for key, c := range store.realationships[name] {
		users = append(users, key)
		chats = append(chats, c)
	}

	delete(store.realationships, name)

	for _, user := range users {
		delete(store.realationships[user], name)
	}
	for _, chat := range chats {
		delete(store.chats, chat)
	}
}

func NewStore() *Store {
	users := make(map[string]*domain.User)
	chats := make(map[string]*domain.Chat)
	realationships := make(map[string]map[string]string)

	return &Store{
		users:          users,
		chats:          chats,
		realationships: realationships,
	}
}
