package redis

import "github.com/matg94/go-url-shortener/errorhandling"

type RedisCache struct {
	cache map[string]string
}

func CreateRedisCache() *RedisCache {
	return &RedisCache{
		cache: make(map[string]string),
	}
}

func (r *RedisCache) GET(key string) (string, error) {
	if r.cache[key] == "" {
		errorhandling.HandleError(ErrRedisValueNotFound, "Cache GET", key)
		return "", ErrRedisValueNotFound
	}
	return r.cache[key], nil
}

func (r *RedisCache) SET(key, value string) error {
	r.cache[key] = value
	return nil
}
