package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	Ctx         = context.TODO()
)
