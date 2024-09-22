package store

import (
	"first_socket/internal/domain"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresStore struct {
	db *gorm.DB
}

func (store *postgresStore) GetUserByLogin(login string) (domain.User, error) {
	var user domain.User
	if err := store.db.Where("login = ?", login).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (store *postgresStore) SaveUser(user domain.User) error {
	if err := store.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func newPostgresStore() (*postgresStore, error) {

	dbHost := os.Getenv("PGHOST")
	dbUser := os.Getenv("PGUSER")
	dbPassword := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGDATABASE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s ", dbHost, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&domain.User{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	store := &postgresStore{
		db: db,
	}

	return store, nil
}
