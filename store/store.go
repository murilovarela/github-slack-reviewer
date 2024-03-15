// writes and reads from the store
package store

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var secrets struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

//encore:service
type StoreService struct {
	db *gorm.DB
}

func NewStoreService() (*StoreService, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", secrets.DBHost, secrets.DBUser, secrets.DBPassword, secrets.DBName, secrets.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return &StoreService{
		db: db,
	}, err
}

// func (s *storeService) Get() {}

// func (s *storeService) Set() {}

// func (s *storeService) Delete() {}

// func (s *storeService) Update() {}

// func (s *storeService) List() {}
