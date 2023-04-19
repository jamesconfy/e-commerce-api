package repo

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache interface {
	Add(key string, value interface{}, expiresAt ...time.Time) error
	AddByTag(key string, value interface{}, tag string, expiry ...time.Time) error
	Get(key string, result interface{}) error
	Delete(key string) error
	DeleteByTag(tag string) error
}

var _ Cache = &redisCache{}

type redisCache struct {
	client *redis.Client
}

func (r *redisCache) Add(key string, value interface{}, expiresAt ...time.Time) error {
	expires := time.Hour * 7

	if expiresAt != nil {
		expires = time.Duration(time.Until(expiresAt[0]))
	}

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(r.client.Context(), key, data, expires).Err()
}

func (r *redisCache) AddByTag(key string, value interface{}, tag string, expiry ...time.Time) error {
	expires := time.Hour * 7

	if expiry != nil {
		expires = time.Duration(time.Until(expiry[0]))
	}

	pipe := r.client.TxPipeline()

	err := pipe.SAdd(r.client.Context(), tag, key).Err()
	if err != nil {
		return err
	}

	err = pipe.Expire(r.client.Context(), tag, expires).Err()
	if err != nil {
		return err
	}

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = pipe.Set(r.client.Context(), key, data, expires).Err()
	if err != nil {
		return err
	}

	_, err = pipe.Exec(r.client.Context())
	return err
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

func (r *redisCache) DeleteByTag(tag string) error {
	keys := make([]string, 0)

	k, _ := r.client.SMembers(r.client.Context(), tag).Result()
	keys = append(keys, tag)
	keys = append(keys, k...)

	return r.client.Del(r.client.Context(), keys...).Err()
}

func NewRedisCache(client *redis.Client) Cache {
	return &redisCache{client: client}
}
