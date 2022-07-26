package services

import (
	"github.com/matg94/go-url-shortener/models"
	"github.com/matg94/go-url-shortener/redis"
	"github.com/matg94/go-url-shortener/repos"
	"github.com/matg94/go-url-shortener/util"
)

func ShortenURL(URLRepo repos.URLRepoInterface, longURL string, hashLength int) (string, error) {
	hash := util.HashString(longURL, hashLength)

	_, err := URLRepo.GetURL(hash)
	if err == redis.ErrRedisValueNotFound {
		err := URLRepo.StoreURL(hash, longURL, 0)
		if err != nil {
			return "", err
		}
		return hash, err
	} else if err != nil {
		return "", err
	}
	return hash, err
}

func ElongateURL(URLRepo repos.URLRepoInterface, shortURL string) (string, error) {
	url, err := URLRepo.GetURL(shortURL)
	if (url == models.URL{}) || err != nil {
		return "", err
	}
	err = URLRepo.IncrementHits(shortURL)
	if err != nil {
		return "", err
	}
	return url.LongURL, nil
}
