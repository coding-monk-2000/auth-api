package storage

import "github.com/coding-monk-2000/auth-api/models"

type AuthStore interface {
	Register(models.User) error
	GetUser(models.Credentials) (*models.User, error)
}
