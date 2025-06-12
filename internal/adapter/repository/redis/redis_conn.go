package redis

import (
	"github.com/redis/go-redis/v9"
)

func RedisConnect() *redis.Client {
	// Connect to Redis
	return redis.NewClient(&redis.Options{
		Addr:     "go_redis:6379", // Redis server address
		Password: "",              // No password
		DB:       0,               // Default DB
	})
}
