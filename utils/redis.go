package utils

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "pass",           // No password set
		DB:       0,                // Use default DB
	})
}

func AddToBlacklist(token string) error {
	expiration := 24 * time.Hour // Set expiration time for the blacklist entry
	return redisClient.Set(ctx, token, "blacklisted", expiration).Err()
}

func IsBlacklisted(token string) (bool, error) {
	result, err := redisClient.Get(ctx, token).Result()
	if err == redis.Nil {
		return false, nil
	}
	return result == "blacklisted", err
}
