package dto

import entity "go-api/internal/core/domain"

type UserCreateRequestBody struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

type UserUpdateRequestBody struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

type UserPasswordRequestBody struct {
	New     string `json:"new"`
	Current string `json:"current"`
}

type UserRequestParam struct {
	ID       uint64 `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

func (u *UserCreateRequestBody) ParseFromEntities(user entity.User) *UserCreateRequestBody {
	return &UserCreateRequestBody{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
}

func (u *UserUpdateRequestBody) ParseFromEntities(user entity.User) *UserUpdateRequestBody {
	return &UserUpdateRequestBody{
		Username: user.Username,
		Email:    user.Email,
	}
}

func (u *UserPasswordRequestBody) ParseFromEntities(password entity.Password) *UserPasswordRequestBody {
	return &UserPasswordRequestBody{
		New:     password.New,
		Current: password.Current,
	}
}

func (u *UserRequestParam) ParseFromEntities(user entity.User) *UserRequestParam {
	return &UserRequestParam{
		ID:       user.ID,
		Username: user.Username,
	}
}
