package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type UserUsecase interface {
	GetUsers(username string) ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserById(ID uint64) (*model.User, error)
	DeleteUser(id_User uint64) error
	UpdateUser(id_User uint64, user model.User) error
	FetchPassword(userID uint64) (string, error)
	UpdatePassword(userID uint64, password string) error
	GetUserByEmail(email string) (model.User, error)
}

type UserUsecaseImpl struct {
	repository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		repository: repo,
	}
}

func (uu UserUsecaseImpl) GetUsers(username string) ([]model.User, error) {
	return uu.repository.GetUsers(username)
}

func (uu UserUsecaseImpl) CreateUser(user model.User) (model.User, error) {
	userID, err := uu.repository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}

	user.ID = userID
	return user, nil
}

func (uu UserUsecaseImpl) GetUserById(ID uint64) (*model.User, error) {
	user, err := uu.repository.GetUserById(ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu UserUsecaseImpl) DeleteUser(id_User uint64) error {
	if err := uu.repository.DeleteUser((id_User)); err != nil {
		return err
	}

	return nil
}

func (uu UserUsecaseImpl) UpdateUser(id_User uint64, user model.User) error {
	if err := uu.repository.UpdateUser(id_User, user); err != nil {
		return err
	}

	return nil
}

func (uu UserUsecaseImpl) FetchPassword(userID uint64) (string, error) {
	passwordInDatabase, err := uu.repository.FetchPassword(userID)
	if err != nil {
		return "", err
	}

	return passwordInDatabase, nil
}

func (uu UserUsecaseImpl) UpdatePassword(userID uint64, password string) error {
	if err := uu.repository.UpdatePassword(userID, password); err != nil {
		return err
	}

	return nil
}

func (uu UserUsecaseImpl) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	user, err := uu.repository.GetUserByEmail(email)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
