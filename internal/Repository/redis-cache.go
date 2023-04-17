package repo

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache interface {
	Add(key string, value interface{}) error
	Get(key string, result interface{}) error
	Delete(key string) error
	// DeleteByTag(tags ...string)
}

var _ Cache = &redisCache{}

type redisCache struct {
	client *redis.Client
}

func (r *redisCache) Add(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(r.client.Context(), key, data, time.Duration(time.Hour*24)).Err()
}

func (r *redisCache) Get(key string, result interface{}) error {
	data, err := r.client.Get(r.client.Context(), key).Bytes()
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return err
	}

	return nil
}

func (r *redisCache) Delete(key string) error {
	return r.client.Del(r.client.Context(), key).Err()
}

func NewRedisCache(client *redis.Client) Cache {
	return &redisCache{client: client}
}
