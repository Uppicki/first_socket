package repositories

import "first_socket/internal/store"

type chatRepository struct {
	*BaseRepository
}

func (repo *chatRepository) RemoveUserChats(name string) {
	repo.store.RemoveUserChats(name)
}

func NewChatRepository(store *store.Store) *chatRepository {
	return &chatRepository{
		BaseRepository: NewBaseRepository(store),
	}
}
