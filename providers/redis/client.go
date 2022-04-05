package redis

import (
	"fmt"
	"quik/config"
	"time"

	"github.com/go-redis/redis"
)

var cache = redis.NewClient(&redis.Options{
	Addr: fmt.Sprintf("%s", config.C.RedisURL),
})

const KEY_DOES_EXIST = redis.Nil

type Redis struct {
}

type RedisInterface interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (int64, error)
	Delete(key string) (error)
}

func (r *Redis) Set(key string, value interface{}, exp time.Duration) error {
	return cache.Set(key, value, exp).Err()
}

func (r *Redis) Get(key string) (int64, error) {
	return cache.Get(key).Int64()
}

func (r *Redis) Delete(key string) (error) {
	return cache.Del(key).Err()
}