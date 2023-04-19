package db

import (
	redis "github.com/go-redis/redis/v8"
)

func NewRedisDB(username, password, address string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Username: username,
		Password: password,
		Addr:     address,
		Network:  "tcp",
		DB:       0,
	})

	return rdb
}
