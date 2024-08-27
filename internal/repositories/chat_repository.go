package repositories

import "first_socket/internal/store"

type ChatRepository struct {
	*BaseRepository
}

func NewChatRepository(store *store.Store) *ChatRepository {
	return &ChatRepository{
		BaseRepository: NewBaseRepository(store),
	}
}
