package storage

import "github.com/redis/go-redis/v9"

func InitRedisClient(address string) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
}

func GetRedisClient() *redis.Client {
	return redisClient
}
