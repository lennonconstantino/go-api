package usecase

import (
	"go-api/internal/adapter/repository"
	entity "go-api/internal/core/domain"
)

type UserUsecase interface {
	GetUsers(username string) ([]entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	GetUserById(userId uint64) (*entity.User, error)
	DeleteUser(userId uint64) error
	UpdateUser(userId uint64, user entity.User) error
	FetchPassword(userId uint64) (string, error)
	UpdatePassword(userId uint64, password string) error
	GetUserByEmail(email string) (entity.User, error)
}

type UserUsecaseImpl struct {
	repository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		repository: repo,
	}
}

func (uu UserUsecaseImpl) GetUsers(username string) ([]entity.User, error) {
	return uu.repository.GetUsers(username)
}

func (uu UserUsecaseImpl) CreateUser(user entity.User) (entity.User, error) {
	userId, err := uu.repository.CreateUser(user)
	if err != nil {
		return entity.User{}, err
	}

	user.ID = userId
	return user, nil
}

func (uu UserUsecaseImpl) GetUserById(userId uint64) (*entity.User, error) {
	return uu.repository.GetUserById(userId)
}

func (uu UserUsecaseImpl) DeleteUser(userId uint64) error {
	return uu.repository.DeleteUser((userId))
}

func (uu UserUsecaseImpl) UpdateUser(userId uint64, user entity.User) error {
	return uu.repository.UpdateUser(userId, user)
}

func (uu UserUsecaseImpl) FetchPassword(userId uint64) (string, error) {
	return uu.repository.FetchPassword(userId)
}

func (uu UserUsecaseImpl) UpdatePassword(userId uint64, password string) error {
	return uu.repository.UpdatePassword(userId, password)
}

func (uu UserUsecaseImpl) GetUserByEmail(email string) (entity.User, error) {
	return uu.repository.GetUserByEmail(email)
}
