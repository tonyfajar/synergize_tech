package db

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient is a singleton Redis client instance
var RedisClient *redis.Client

// InitRedis initializes the Redis client
func InitRedis() {
	options := &redis.Options{
		Addr:     "localhost:6379", // Replace with your Redis server address
		Password: "",               // No password by default
		DB:       0,                // Default DB
	}

	RedisClient = redis.NewClient(options)
}

// SetKey sets a key-value pair in Redis
func SetKey(key, value string, expiration time.Duration) error {
	return RedisClient.Set(context.Background(), key, value, expiration).Err()
}

// GetKey gets the value for a key from Redis
func GetKey(key string) (string, error) {
	return RedisClient.Get(context.Background(), key).Result()
}

// CloseRedis closes the Redis client
func CloseRedis() error {
	return RedisClient.Close()
}
