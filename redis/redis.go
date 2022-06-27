package redis

import (
	"errors"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/matg94/go-url-shortener/config"
)

type RedisConnectionInterface interface {
	GET(key string) (string, error)
	SET(key, value string) error
	DEL(keys ...string) error
}

var ErrSomeRedisError error = errors.New("some redis error")

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
			if redisConfig.Password == "" {
				c, err = redis.Dial(
					"tcp",
					fmt.Sprintf(
						"%s:%d",
						redisConfig.URL,
						redisConfig.Port,
					),
				)
			} else {
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
			}
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
	return "", nil
}

func (r *RedisConnection) SET(key, value string) error {
	return nil
}

func (r *RedisConnection) DEL(key string) error {
	return nil
}
