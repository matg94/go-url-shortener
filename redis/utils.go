package redis

import (
	"log"

	"github.com/matg94/go-url-shortener/config"
)

func SetupDataPersistence(redisConfig config.RedisConfig) RedisConnectionInterface {
	if redisConfig.IsCache {
		log.Print("Starting data persistence with CACHE")
		return CreateRedisCache()
	}
	log.Printf("Starting data persistence with redis at %s:%d", redisConfig.URL, redisConfig.Port)
	return CreateRedisConnectionPool(&redisConfig)
}
