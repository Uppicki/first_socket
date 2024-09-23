package store

import "first_socket/internal/domain"

type IStore interface {
	Migrate() error
	GetUserByLogin(string) (domain.User, error)
	SaveUser(domain.User) error
}

func NewStore(typeStore string) (IStore, error) {
	var store IStore
	var err error

	switch typeStore {
	case "postgres":
		store, err = newPostgresStore()
	}

	return store, err
}
