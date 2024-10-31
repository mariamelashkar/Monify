package redis

import (
	"context"
	"os"
	"github.com/go-redis/redis/v8"
	//"log"
)

var Ctx = context.Background()
var Rdb *redis.Client

func InitRedis() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   0, // use default DB
	})

	// // Test the connection
	// _, err := Rdb.Ping(Ctx).Result()
	// if err != nil {
	// 	log.Fatalf("Could not connect to Redis: %v", err)
	// } else {
	// 	log.Println("Connected to Redis successfully")
	// }
}
