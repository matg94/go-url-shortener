package repos

import (
	"github.com/matg94/go-url-shortener/models"
	"github.com/matg94/go-url-shortener/redis"
)

type URLRepoInterface interface {
	StoreURL(shortURL string, newURL models.URL) error
	GetURL(shortURL string) (models.URL, error)
	IncrementHits(shortURL string) error
}

type URLRepo struct {
	RedisConn redis.RedisConnectionInterface
}

func CreateURLRepo(redisConnection redis.RedisConnectionInterface) URLRepoInterface {
	return &URLRepo{
		RedisConn: redisConnection,
	}
}

// Should this be changed to take URL model as a param?
func (repo *URLRepo) StoreURL(shortURL string, newURL models.URL) error {
	url := models.URL{
		LongURL: newURL.LongURL,
		Hits:    newURL.Hits,
	}
	url_json, err := url.ToJSON()
	if err != nil {
		return err
	}
	return repo.RedisConn.SET(shortURL, url_json)
}

func (repo *URLRepo) GetURL(shortURL string) (models.URL, error) {
	url_data, err := repo.RedisConn.GET(shortURL)
	if err != nil {
		return models.URL{}, err // Should return an error if not found that can be interpreted further up
	}
	url, err := models.FromJSON(url_data)
	return url, err
}

func (repo *URLRepo) IncrementHits(shortURL string) error {
	url, err := repo.GetURL(shortURL)
	if err != nil {
		return err
	}
	url.Hits += 1
	err = repo.StoreURL(shortURL, url)
	if err != nil {
		return err
	}
	return nil
}
