package services

import (
	"github.com/matg94/go-url-shortener/redis"
	"github.com/matg94/go-url-shortener/util"
)

func ShortenURL(redisConn redis.RedisConnectionInterface, longURL string, hashLength int) (string, error) {
	hash := util.HashString(longURL, hashLength)

	foundHash, err := redisConn.GET(hash)
	if err != nil {
		return "", err
	}
	if foundHash == "" {
		err = redisConn.SET(hash, longURL)
		if err != nil {
			return "", err
		}
	}
	return hash, err
}

func ElongateURL(redisConn redis.RedisConnectionInterface, shortURL string) (string, error) {
	foundURL, err := redisConn.GET(shortURL)
	if err != nil || foundURL == "" {
		return "", err
	}

	return foundURL, nil
}
