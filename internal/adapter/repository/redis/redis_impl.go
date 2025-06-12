package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// CacheRepository signature of methods
type CacheRepository interface {
	Set(key string, value []byte, duration time.Duration) error
	Get(key string) ([]byte, error)
	Delete(key string) (int64, error)
	FlushDB() error
	//HandleError(err error)
}

// CacheRepositoryImpl implements the methods
type CacheRepositoryImpl struct {
	connect *redis.Client
}

func NewCacheRepository(connect *redis.Client) *CacheRepositoryImpl {
	return &CacheRepositoryImpl{
		connect: connect,
	}
}

func (c CacheRepositoryImpl) Set(key string, value []byte, duration *time.Duration) error {
	//expiration, _ := time.ParseDuration("30us")
	// hr, _ := time.ParseDuration("5h")
	// comp, _ := time.ParseDuration("2h30m40s")
	// m1, _ := time.ParseDuration("3µs")
	// m2, _ := time.ParseDuration("3us")
	var expiration time.Duration
	if duration == nil {
		expiration = 10 * time.Second
	} else {
		expiration = *duration
	}
	var ctx = context.Background()
	// Set and Get a key-value pair
	err := c.connect.Set(ctx, key, value, expiration).Err()
	if err != nil {
		fmt.Println("Failed to set key")
		return err
	}
	return nil
}

func (c CacheRepositoryImpl) Get(key string) ([]byte, error) {
	var ctx = context.Background()

	val, err := c.connect.Get(ctx, key).Bytes()
	if err != nil {
		fmt.Println("Failed to get key")
		return nil, err
	}
	return val, nil
}

func (c CacheRepositoryImpl) Delete(key string) (int64, error) {
	var ctx = context.Background()

	deletedCount, err := c.connect.Del(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if deletedCount > 0 {
		fmt.Println("Key deleted successfully:", key)
	} else {
		fmt.Println("Key not found:", key)
	}

	return deletedCount, nil
}

func (c CacheRepositoryImpl) FlushDB() error {
	var ctx = context.Background()

	err := c.connect.FlushDB(ctx).Err()
	if err != nil {
		fmt.Println("Failed to Flush")
		return err
	}
	return nil
}

// func (c CacheRepositoryImpl) HandleError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
