package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func RedisConnect() *redis.Client {
	// Connect to Redis
	return redis.NewClient(&redis.Options{
		Addr:     "go_redis:6379", // Redis server address
		Password: "",              // No password
		DB:       0,               // Default DB
	})
}

func Set(key string, value []byte) error {
	c := RedisConnect()
	defer c.Close()

	//expiration, _ := time.ParseDuration("30us")
	// hr, _ := time.ParseDuration("5h")
	// comp, _ := time.ParseDuration("2h30m40s")
	// m1, _ := time.ParseDuration("3µs")
	// m2, _ := time.ParseDuration("3us")
	expiration := 10 * time.Second

	// Set and Get a key-value pair
	err := c.Set(ctx, key, value, expiration).Err()
	if err != nil {
		//HandleError(fmt.Errorf("Failed to set key"))
		fmt.Println("Failed to set key")
		return err
	}
	return nil
}

func Get(key string) ([]byte, error) {
	c := RedisConnect()
	defer c.Close()

	//c.Expire(context.Background(), key, 1*time.Second)

	val, err := c.Get(ctx, key).Bytes()
	if err != nil {
		//HandleError(fmt.Errorf("Failed to get key"))
		fmt.Println("Failed to get key")
		return nil, err
	}
	return val, nil
}

func Flush(key string) error {
	c := RedisConnect()
	defer c.Close()

	err := c.FlushDB(ctx).Err()
	if err != nil {
		//HandleError(fmt.Errorf("Failed to Flush"))
		fmt.Println("Failed to Flush")
		return err
	}
	return nil
}
