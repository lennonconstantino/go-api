package model

import (
	"errors"
	"go-api/utils"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitzero"`
}

// Prepare will call the methods to validate and format the received user
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

// validate used to validate the fields of the user structure
func (user *User) validate(step string) error {
	if user.Username == "" {
		return errors.New("The name is required and cannot be blank.")
	}
	if user.Email == "" {
		return errors.New("Email is required and cannot be blank.")
	}
	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("The email entered is invalid")
	}
	if step == "form" && user.Password == "" {
		return errors.New("Password is mandatory and cannot be blank.")
	}

	return nil
}

// format used to remove white spaces
func (user *User) format(step string) error {
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)

	if step == "form" {
		passwordWithHash, err := utils.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passwordWithHash)
	}
	return nil
}
