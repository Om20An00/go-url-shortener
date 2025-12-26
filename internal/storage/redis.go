package storage

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var RedisCtx = context.Background()

func ConnectRedis() {
	// Use REDIS_ADDR from environment (Docker-friendly)
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379" // fallback for local development
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // leave empty if no password
		DB:       0,
	})

	_, err := Rdb.Ping(RedisCtx).Result()
	if err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	log.Println("Connected to Redis")
}
