package repository

import (
	entity "go-api/internal/core/domain"
)

// UserRespoistory signature of methods
type UserRepository interface {
	GetUsers(username string) ([]entity.User, error)
	CreateUser(user entity.User) (uint64, error)
	GetUserById(userId uint64) (*entity.User, error)
	DeleteUser(userId uint64) error
	UpdateUser(userId uint64, user entity.User) error
	FetchPassword(userId uint64) (string, error)
	UpdatePassword(userId uint64, password string) error
	GetUserByEmail(email string) (entity.User, error)
}
