package postgres

import (
	"database/sql"
	"fmt"
	entity "go-api/internal/core/domain"
)

// UserRepositoryImpl implements the methods
type UserRepositoryImpl struct {
	connection *sql.DB
}

// NewUserRepository initialize the repo
func NewUserRepository(connection *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		connection: connection,
	}
}

// GetUser
func (ur UserRepositoryImpl) GetUsers(username string) ([]entity.User, error) {
	username = fmt.Sprintf("%%%s%%", username) // %username%

	rows, err := ur.connection.Query("select id, username, email, created_at from users where username LIKE $1", username)
	if err != nil {
		fmt.Println(err)
		return []entity.User{}, err
	}
	defer rows.Close()

	var UserList []entity.User
	var UserObj entity.User

	for rows.Next() {
		err = rows.Scan(
			&UserObj.ID,
			&UserObj.Username,
			&UserObj.Email,
			&UserObj.CreatedAt)

		if err != nil {
			fmt.Println(err)
			return []entity.User{}, err
		}

		UserList = append(UserList, UserObj)
	}

	rows.Close()

	return UserList, nil
}

// CreateUser
func (ur UserRepositoryImpl) CreateUser(user entity.User) (uint64, error) {
	var id int
	query, err := ur.connection.Prepare("INSERT INTO users" +
		"(username, email, password)" +
		" VALUES ($1, $2, $3) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return uint64(id), nil
}

// GetUserById
func (ur UserRepositoryImpl) GetUserById(userId uint64) (*entity.User, error) {
	query, err := ur.connection.Prepare("SELECT id, username, email, created_at FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer query.Close()

	var user entity.User

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

	return &user, nil
}

// DeleteUser
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

// UpdateUser
func (ur UserRepositoryImpl) UpdateUser(userId uint64, user entity.User) error {
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

	var user entity.User

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

// GetUserByEmail
func (ur UserRepositoryImpl) GetUserByEmail(email string) (entity.User, error) {
	line, err := ur.connection.Query("SELECT id, password FROM users WHERE email = $1", email)
	if err != nil {
		fmt.Println(err)
		return entity.User{}, err
	}
	defer line.Close()

	var user entity.User
	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return entity.User{}, err
		}
	}

	return user, nil
}
