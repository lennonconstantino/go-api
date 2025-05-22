package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type UserRepository interface {
	GetUsers(username string) ([]model.User, error)
	CreateUser(user model.User) (uint64, error)
	GetUserById(userId uint64) (*model.User, error)
	DeleteUser(userId uint64) error
	UpdateUser(userId uint64, user model.User) error
	FetchPassword(userId uint64) (string, error)
	UpdatePassword(userId uint64, password string) error
	GetUserByEmail(email string) (model.User, error)
}

type UserRepositoryImpl struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		connection: connection,
	}
}

func (ur UserRepositoryImpl) GetUsers(username string) ([]model.User, error) {
	username = fmt.Sprintf("%%%s%%", username) // %username%

	rows, err := ur.connection.Query("select id, username, email, created_at from users where username LIKE $1", username)
	if err != nil {
		fmt.Println(err)
		return []model.User{}, err
	}

	var UserList []model.User
	var UserObj model.User

	for rows.Next() {
		err = rows.Scan(
			&UserObj.ID,
			&UserObj.Username,
			&UserObj.Email,
			&UserObj.CreatedAt)

		if err != nil {
			fmt.Println(err)
			return []model.User{}, err
		}

		UserList = append(UserList, UserObj)
	}

	rows.Close()

	return UserList, nil
}

func (ur UserRepositoryImpl) CreateUser(user model.User) (uint64, error) {
	var id int
	query, err := ur.connection.Prepare("INSERT INTO users" +
		"(username, email, password)" +
		" VALUES ($1, $2, $3) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return uint64(id), nil
}

func (ur UserRepositoryImpl) GetUserById(userId uint64) (*model.User, error) {
	query, err := ur.connection.Prepare("SELECT id, username, email, created_at FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user model.User

	err = query.QueryRow(userId).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &user, nil
}

func (ur UserRepositoryImpl) DeleteUser(userId uint64) error {
	statement, err := ur.connection.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId); err != nil {
		return err
	}

	return nil
}

func (ur UserRepositoryImpl) UpdateUser(userId uint64, user model.User) error {
	statement, err := ur.connection.Prepare(
		"update users set username = $1, email = $2 where id = $3",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Username, user.Email, userId); err != nil {
		return err
	}

	return nil
}

// FetchPassword fetches a user's password by ID
func (ur UserRepositoryImpl) FetchPassword(userId uint64) (string, error) {
	line, err := ur.connection.Query("select password from users where id = $1", userId)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user model.User

	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

// UpdatePassword changes the password of a user
func (ur UserRepositoryImpl) UpdatePassword(userId uint64, password string) error {
	statement, err := ur.connection.Prepare("update users set password = $1 where id = $2")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userId); err != nil {
		return err
	}

	return nil
}

func (ur UserRepositoryImpl) GetUserByEmail(email string) (model.User, error) {
	line, err := ur.connection.Query("SELECT id, password FROM users WHERE email = $1", email)
	if err != nil {
		fmt.Println(err)
		return model.User{}, err
	}
	defer line.Close()

	var user model.User
	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}
