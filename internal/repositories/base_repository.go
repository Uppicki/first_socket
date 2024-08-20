package repositories

import "first_socket/internal/store"

type BaseRepository struct {
	store *store.Store
}

func NewBaseRepository(store *store.Store) *BaseRepository {
	return &BaseRepository{
		store: store,
	}
}
