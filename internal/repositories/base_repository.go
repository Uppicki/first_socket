package repositories

import "first_socket/internal/store"

type BaseRepository struct {
	store store.IStore
}

func NewBaseRepository(store store.IStore) *BaseRepository {
	return &BaseRepository{
		store: store,
	}
}
