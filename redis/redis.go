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
	conn := r.RedisPool.Get()
	val, err := conn.Do("GET", key)
	defer conn.Close()

	if err != nil {
		errorhandling.HandleError(err, "Redis GET on key", key)
		return "", err
	}
	if val == "" {
		return "", ErrRedisValueNotFound
	}
	return fmt.Sprint(val), nil
}

func (r *RedisConnection) SET(key, value string) error {
	conn := r.RedisPool.Get()
	_, err := conn.Do("SET", key, value)
	defer conn.Close()

	if err != nil {
		errorhandling.HandleError(err, "Redis SET on key, value", fmt.Sprint(key, value))
		return err
	}
	return nil
}
