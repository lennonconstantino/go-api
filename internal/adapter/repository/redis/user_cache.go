package redis

import (
	"encoding/json"
	"fmt"
	"go-api/internal/adapter/repository"
	entity "go-api/internal/core/domain"
)

// UserCacheRespoistory signature of methods
type UserCacheRepository interface {
	GetUsers(username string) ([]entity.User, error)
	GetUserById(userId uint64) (*entity.User, error)
}

// UserCacheRepositoryImpl implements the methods
type UserCacheRepositoryImpl struct {
	cache CacheRepository
	user  repository.UserRepository
}

// NewUserCacheRepository initialize the repo
func NewUserCacheRepository(cache CacheRepository, user repository.UserRepository) *UserCacheRepositoryImpl {
	return &UserCacheRepositoryImpl{
		cache: cache,
		user:  user,
	}
}

func (ur UserCacheRepositoryImpl) GetUsers(username string) ([]entity.User, error) {
	var users []entity.User

	reply, err := ur.cache.Get(fmt.Sprintf("users:%s", username))
	if err != nil {
		users, err = ur.user.GetUsers(username)
		if err != nil {
			return nil, err
		}
		userBytes, _ := json.Marshal(users)
		ur.cache.Set(fmt.Sprintf("users:%s", username), userBytes, 10)

		return users, nil
	}

	json.Unmarshal(reply, &users)
	return users, nil
}

func (ur UserCacheRepositoryImpl) GetUserById(userId uint64) (*entity.User, error) {
	var user *entity.User

	reply, err := ur.cache.Get(fmt.Sprintf("users:%d", userId))
	if err != nil {
		user, err = ur.user.GetUserById(userId)
		if err != nil {
			return nil, err
		}
		userBytes, _ := json.Marshal(user)
		ur.cache.Set(fmt.Sprintf("users:%d", userId), userBytes, 10)

		return user, nil
	}

	json.Unmarshal(reply, &user)
	return user, nil
}
