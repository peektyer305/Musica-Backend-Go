package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
)

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}
