package cache

import (
	"context"
	"log"
	"matrimonial-service/internal/config"

	"github.com/redis/go-redis/v9"
)

// NewRedisClient creates and returns a new Redis client
func NewRedisClient(cfg config.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// Ping to check connection
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Println("Redis connection established successfully.")
	return rdb
}
