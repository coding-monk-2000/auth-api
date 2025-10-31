package storage

import (
	"github.com/coding-monk-2000/auth-api/models"
	"gorm.io/gorm"
)

type GormStore struct {
	DB *gorm.DB
}

func NewGormStore(db *gorm.DB) *GormStore {
	db.AutoMigrate(&models.User{})
	return &GormStore{DB: db}
}

func (s *GormStore) Register(user models.User) error {
	return s.DB.Create(&user).Error
}

func (s *GormStore) GetUser(creds models.Credentials) (*models.User, error) {
	var user models.User
	err := s.DB.Where("username = ?", creds.Username).First(&user).Error
	return &user, err
}
