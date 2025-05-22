package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type UserUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return UserUsecase{
		repository: repo,
	}
}

func (uu *UserUsecase) GetUsers(username string) ([]model.User, error) {
	return uu.repository.GetUsers(username)
}

func (uu *UserUsecase) CreateUser(user model.User) (model.User, error) {
	userID, err := uu.repository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}

	user.ID = userID
	return user, nil
}

func (uu *UserUsecase) GetUserById(ID uint64) (*model.User, error) {
	user, err := uu.repository.GetUserById(ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *UserUsecase) DeleteUser(id_User uint64) error {
	if err := uu.repository.DeleteUser((id_User)); err != nil {
		return err
	}

	return nil
}

func (uu *UserUsecase) UpdateUser(id_User uint64, user model.User) error {
	if err := uu.repository.UpdateUser(id_User, user); err != nil {
		return err
	}

	return nil
}

func (uu *UserUsecase) FetchPassword(userID uint64) (string, error) {
	passwordInDatabase, err := uu.repository.FetchPassword(userID)
	if err != nil {
		return "", err
	}

	return passwordInDatabase, nil
}

func (uu *UserUsecase) UpdatePassword(userID uint64, password string) error {
	if err := uu.repository.UpdatePassword(userID, password); err != nil {
		return err
	}

	return nil
}

func (uu *UserUsecase) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	user, err := uu.repository.GetUserByEmail(email)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
