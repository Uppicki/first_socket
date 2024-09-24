package repositories

import (
	"first_socket/internal/domain"
	"first_socket/internal/store"
)

type UserRepository struct {
	*BaseRepository
}

func (repo *UserRepository) IsLoginExsist(login string) bool {
	_, err := repo.store.GetUserByLogin(login)

	return err != nil
}

func (repo *UserRepository) CreateUser(user domain.User) error {
	err := repo.store.SaveUser(user)

	return err
}

func (repo *UserRepository) GetUserByLogin(login string) (domain.User, error) {
	user, err := repo.store.GetUserByLogin(login)

	return user, err
}

func NewUserRepository(store store.IStore) IUserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository(store),
	}
}
