package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Redis struct
type Redis struct {
	client      *redis.Client
	cacheExpire time.Duration
}

// NewCacheRedis Redis's Constructor
func NewCacheRedis(client *redis.Client, cacheExpire time.Duration) *Redis {
	return &Redis{client: client, cacheExpire: cacheExpire}
}

// Get method
func (r *Redis) Get(key string, value interface{}) Output {
	res, err := r.client.Get(key).Result()
	if err != nil {
		return Output{Result: nil, Error: err}
	}

	var jsonBlob = []byte(res)
	err = json.Unmarshal(jsonBlob, value)
	if err != nil {
		return Output{Result: nil, Error: err}
	}

	return Output{Result: value, Error: nil}
}

// Set method
func (r *Redis) Set(key string, data interface{}) Output {
	payload, err := json.Marshal(data)
	if err != nil {
		return Output{Result: nil, Error: err}
	}

	err = r.client.Set(key, payload, r.cacheExpire).Err()
	if err != nil {
		return Output{Result: nil, Error: err}
	}
	return Output{Result: nil, Error: nil}
}

// Del method
func (r *Redis) Del(key string) Output {
	key = fmt.Sprintf("*%s*", key)
	vals, err := r.client.Keys(key).Result()
	if err != nil {
		return Output{Result: nil, Error: err}
	}

	for _, v := range vals {
		err = r.client.Del(v).Err()
		if err != nil {
			return Output{Result: nil, Error: err}
		}
	}
	return Output{Result: nil, Error: nil}
}

// Flush method
func (r *Redis) Flush() Output {
	res, err := r.client.FlushDB().Result()
	return Output{Result: res, Error: err}
}
