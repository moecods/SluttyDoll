package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6800",
		//Password: "123",
		//DB:       0,
	})

	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Println("Successfully connected to Redis!")
}
