package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient  *redis.Client
	RedisContext = context.TODO()
)
