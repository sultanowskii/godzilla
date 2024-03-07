package storage

import "github.com/redis/go-redis/v9"

func InitRedisClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func GetRedisClient() *redis.Client {
	return redisClient
}