package redis

import (
	"errors"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/matg94/go-url-shortener/config"
	"github.com/matg94/go-url-shortener/errorhandling"
)

type RedisConnectionInterface interface {
	GET(key string) (string, error)
	SET(key, value string) error
}

var ErrRedisValueNotFound error = errors.New("short url key not found")

type RedisConnection struct {
	RedisPool *redis.Pool
}

func CreateRedisConnectionPool(redisConfig *config.RedisConfig) *RedisConnection {
	pool := &redis.Pool{
		MaxIdle:   redisConfig.MaxIdle,
		MaxActive: redisConfig.MaxActive,
		Dial: func() (redis.Conn, error) {
			var c redis.Conn
			var err error
			c, err = redis.Dial(
				"tcp",
				fmt.Sprintf(
					"%s:%d",
					redisConfig.URL,
					redisConfig.Port,
				),
				redis.DialPassword(redisConfig.Password),
				redis.DialUseTLS(true),
			)
			if err != nil {
				log.Fatal(err.Error())
			}
			return c, err
		},
	}
	return &RedisConnection{
		RedisPool: pool,
	}
}

func (r *RedisConnection) GET(key string) (string, error) {
	err := errors.New("")
	errorhandling.HandleError(err, "Redis GET on key", key)
	return "", nil
}

func (r *RedisConnection) SET(key, value string) error {
	err := errors.New("")
	errorhandling.HandleError(err, "Redis SET on key, value", fmt.Sprint(key, value))
	return nil
}
