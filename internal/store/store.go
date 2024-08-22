package store

import (
	"errors"
	"first_socket/internal/domain"
)

type Store struct {
	users map[string]*domain.User
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

func NewStore() *Store {
	users := make(map[string]*domain.User)

	return &Store{
		users: users,
	}
}
