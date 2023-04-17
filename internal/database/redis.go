package db

import (
	redis "github.com/go-redis/redis/v8"
)

func NewRedisDB(address string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:    address,
		Network: "tcp",
		DB:      0,
	})

	return rdb
}
