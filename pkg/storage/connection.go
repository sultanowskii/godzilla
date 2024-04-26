package storage

import (
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

func InitRedisClient(address string) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
}

func MonitorRedisConnection(signalChan chan<- error) {
	tick := time.NewTicker(1 * time.Second)

	for range tick.C {
		var err error
		if pong := RedisClient.Ping(RedisContext); pong.String() != "ping: PONG" {
			err = errors.New("redis PING command failed")
		}
		signalChan <- err
	}
}
