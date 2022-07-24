package services

import (
	"fmt"
	"testing"

	"github.com/matg94/go-url-shortener/models"
	"github.com/matg94/go-url-shortener/redis"
	"github.com/matg94/go-url-shortener/repos"
)

func TestShortenURL(t *testing.T) {

	redisMock := redis.CreateRedisMock(
		nil,
		nil,
		"",
	)

	longURL := "test"
	hashLength := 10

	url_repo := &repos.URLRepo{
		RedisConn: redisMock,
	}

	shortened, err := ShortenURL(url_repo, longURL, hashLength)
	if err != nil {
		t.Log("expected no errors but got", err)
		t.Fail()
	}

	if shortened == "" {
		t.Log("expected returned shortened url but got empty")
		t.Fail()
	}
}

func TestShortenURLRedisErrorGet(t *testing.T) {
	longURL := "test"
	hashLength := 10
	redisMock := &redis.RedisConnectionMock{}
	redisMock.ReturnValue = ""
	redisError := fmt.Errorf("fail")
	redisMock.ReturnErrorGet = redisError

	url_repo := &repos.URLRepo{
		RedisConn: redisMock,
	}

	shortened, err := ShortenURL(url_repo, longURL, hashLength)
	if err != redisError {
		t.Log("expected error to be nil but got", err)
		t.Fail()
	}

	if shortened != "" {
		t.Log("expected empty string but got ", shortened)
		t.Fail()
	}
}

func TestShortenURLRedisErrorSet(t *testing.T) {
	longURL := "test"
	hashLength := 10
	redisMock := &redis.RedisConnectionMock{}
	redisMock.ReturnValue = ""
	redisError := fmt.Errorf("fail")
	redisMock.ReturnErrorSet = redisError

	url_repo := &repos.URLRepo{
		RedisConn: redisMock,
	}

	shortened, err := ShortenURL(url_repo, longURL, hashLength)
	if err != redisError {
		t.Log("expected error to be nil but got", err)
		t.Fail()
	}

	if shortened != "" {
		t.Log("expected empty string but got ", shortened)
		t.Fail()
	}
}

func TestElongateURL(t *testing.T) {
	shortURL := "test"
	redisMock := &redis.RedisConnectionMock{}
	longURL := &models.URL{
		LongURL: "test-long",
		Hits:    5,
	}
	redisMock.ReturnValue, _ = longURL.ToJSON()

	url_repo := &repos.URLRepo{
		RedisConn: redisMock,
	}

	url, err := ElongateURL(url_repo, shortURL)

	if err != nil {
		t.Log("expected error to be nil but got", err)
		t.Fail()
	}

	if url != "test-long" {
		t.Log("expected url to be 'test-long' but got", url)
		t.Fail()
	}
}

func TestElongateURLRedisErrorGet(t *testing.T) {
	shortURL := "test"
	redisMock := &redis.RedisConnectionMock{}
	redisMock.ReturnValue = "test-long"
	redisError := fmt.Errorf("fail")
	redisMock.ReturnErrorGet = redisError

	url_repo := &repos.URLRepo{
		RedisConn: redisMock,
	}

	url, err := ElongateURL(url_repo, shortURL)

	if err != redisError {
		t.Log("expected error to be 'fail' but got", err)
		t.Fail()
	}

	if url != "" {
		t.Log("expected url to be empty but got", url)
		t.Fail()
	}
}

func TestElongateURLRedisErrorIncrement(t *testing.T) {

}
