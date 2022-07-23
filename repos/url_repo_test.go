package repos

import (
	"testing"

	"github.com/matg94/go-url-shortener/models"
	"github.com/matg94/go-url-shortener/redis"
)

func TestStoreURL(t *testing.T) {
	longURL := "test"
	shortURL := "123123"
	hits := uint(1)

	redisMock := &redis.RedisConnectionMock{}

	repo := &URLRepo{
		redisConn: redisMock,
	}

	err := repo.StoreURL(shortURL, longURL, hits)
	if err != nil {
		t.Logf("expected no errors but got %s", err)
		t.Fail()
	}

	if redisMock.ReturnErrorSet != nil {
		t.Logf("expected no errors from redis but got %s", redisMock.ReturnErrorSet)
		t.Fail()
	}

	url := models.URL{
		LongURL: longURL,
		Hits:    hits,
	}
	url_json, _ := url.ToJSON()

	if redisMock.Value != url_json {
		t.Logf("expected redis value set to be valid json but got %s", url_json)
		t.Fail()
	}

}

func TestShortenURLRedisErrorGet(t *testing.T) {

}

func TestShortenURLRedisErrorSet(t *testing.T) {

}

func TestElongateURL(t *testing.T) {

}

func TestElongateURLRedisError(t *testing.T) {

}
