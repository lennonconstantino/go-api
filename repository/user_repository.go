package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetUsers(username string) ([]model.User, error) {
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

func (ur *UserRepository) CreateUser(User model.User) (uint64, error) {
	var id int
	query, err := ur.connection.Prepare("INSERT INTO users" +
		"(username, email, password)" +
		" VALUES ($1, $2, $3) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(User.Username, User.Email, User.Password).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return uint64(id), nil
}

func (ur *UserRepository) GetUserById(ID uint64) (*model.User, error) {
	query, err := ur.connection.Prepare("SELECT id, username, email, created_at FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var User model.User

	err = query.QueryRow(ID).Scan(
		&User.ID,
		&User.Username,
		&User.Email,
		&User.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &User, nil
}

func (ur *UserRepository) DeleteUser(id_User uint64) error {
	statement, err := ur.connection.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id_User); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) UpdateUser(id_User uint64, User model.User) error {
	statement, err := ur.connection.Prepare(
		"update users set username = $1, email = $2 where id = $3",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(User.Username, User.Email, id_User); err != nil {
		return err
	}

	return nil
}

// FetchPassword fetches a user's password by ID
func (ur *UserRepository) FetchPassword(userID uint64) (string, error) {
	line, err := ur.connection.Query("select password from users where id = $1", userID)
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
func (ur *UserRepository) UpdatePassword(userID uint64, password string) error {
	statement, err := ur.connection.Prepare("update users set password = $1 where id = $2")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userID); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserByEmail(email string) (model.User, error) {
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
