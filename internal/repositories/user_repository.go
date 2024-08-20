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

func NewUserRepository(store *store.Store) *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository(store),
	}
}
