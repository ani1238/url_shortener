package redisdb

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

// InitializeRedisClient initializes the Redis client.
func InitializeRedisClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Use the service name as the hostname
		Password: "",               // No password by default
		DB:       0,                // Default DB
	})

	// Check if the Redis server is reachable.
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}

// ShortenURL stores the shortened URL in Redis.
func AddToRedis(id string, data interface{}, ttl time.Duration) error {
	return redisClient.Set(ctx, id, data, ttl).Err()
}

// GetLongURL retrieves the long URL from Redis.
func GetFromRedis(id string) (string, error) {
	return redisClient.Get(ctx, id).Result()
}

func GetShortURLID(longURL string) (string, error) {
	return redisClient.Get(ctx, "long:"+longURL).Result()
}
