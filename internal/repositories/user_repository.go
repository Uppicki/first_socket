package repositories

import (
	"first_socket/internal/domain"
	"first_socket/internal/store"
)

type UserRepository struct {
	*BaseRepository
}

func (repository *UserRepository) AddUser(name string) {
	repository.store.AddUser(name)
}

func (repository *UserRepository) GetAllUsers() map[string]*domain.User {
	users := repository.store.GetAllUsers()
	return users
}

func (repository *UserRepository) GetUsernamesWithoutUser(
	name string,
) []string {
	users := repository.store.GetAllUsers()

	if _, ok := users[name]; ok {
		delete(users, name)
	}

	usernames := make([]string, 0)

	for key, _ := range users {
		usernames = append(usernames, key)
	}

	return usernames
}

func (repository *UserRepository) GetUserByName(
	name string,
) (*domain.User, error) {
	return repository.store.GetUserByName(name)
}

func NewUserRepository(store *store.IStore) *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository(store),
	}
}
