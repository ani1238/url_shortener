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

const domainCountKey = "domain_counts"

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

// IncrementDomainCount increments the count for a domain in a Redis sorted set.
func IncrementDomainCount(domain string) error {
	return redisClient.ZIncrBy(ctx, "domain_counts", 1, domain).Err()
}

// GetTopDomains retrieves the top N domains with the highest counts from the Redis sorted set.
func GetTopDomains(N int) ([]string, error) {
	topDomains, err := redisClient.ZRevRange(ctx, "domain_counts", 0, int64(N-1)).Result()
	if err != nil {
		return nil, err
	}
	return topDomains, nil
}

// GetDomainCount retrieves the count for a domain from the Redis sorted set.
func GetDomainCount(domain string) (float64, error) {
	return redisClient.ZScore(ctx, "domain_counts", domain).Result()
}
