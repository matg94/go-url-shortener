package services

import (
	"github.com/matg94/go-url-shortener/models"
	"github.com/matg94/go-url-shortener/repos"
	"github.com/matg94/go-url-shortener/util"
)

func ShortenURL(URLRepo *repos.URLRepo, longURL string, hashLength int) (string, error) {
	hash := util.HashString(longURL, hashLength)

	url, err := URLRepo.GetURL(hash)
	if err != nil {
		return "", err
	}

	if (url == models.URL{}) {
		err := URLRepo.StoreURL(hash, longURL, 0)
		if err != nil {
			return "", err
		}
	}

	return hash, err
}

func ElongateURL(URLRepo *repos.URLRepo, shortURL string) (string, error) {
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
