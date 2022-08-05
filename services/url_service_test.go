package services

import (
	"fmt"
	"testing"

	"github.com/matg94/go-url-shortener/models"
	"github.com/matg94/go-url-shortener/redis"
	"github.com/matg94/go-url-shortener/repos"
)

func TestShortenURLFound(t *testing.T) {
	longURL := "test"
	hashLength := 10

	url := &models.URL{
		LongURL: longURL,
		Hits:    1,
	}
	url_json, _ := url.ToJSON()

	redisMock := redis.CreateRedisMock(nil, nil, url_json)
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

func TestShortenURLNotFound(t *testing.T) {

	redisMock := redis.CreateRedisMock(redis.ErrRedisValueNotFound, nil, "")
	url_repo := &repos.URLRepo{
		RedisConn: redisMock,
	}

	longURL := "test"
	hashLength := 10

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

	redisError := fmt.Errorf("fail")
	redisMock := redis.CreateRedisMock(redisError, nil, "")
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

	redisError := fmt.Errorf("fail")
	redisMock := redis.CreateRedisMock(redis.ErrRedisValueNotFound, redisError, "")
	url_repo := &repos.URLRepo{
		RedisConn: redisMock,
	}

	shortened, err := ShortenURL(url_repo, longURL, hashLength)
	if err != redisError {
		t.Log("expected error to be fail but got", err)
		t.Fail()
	}

	if shortened != "" {
		t.Log("expected empty string but got ", shortened)
		t.Fail()
	}
}

func TestElongateURL(t *testing.T) {
	shortURL := "test"
	longURL := &models.URL{
		LongURL: "test-long",
		Hits:    5,
	}
	url, _ := longURL.ToJSON()
	redisMock := redis.CreateRedisMock(nil, nil, url)
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
	redisError := fmt.Errorf("fail")
	redisMock := redis.CreateRedisMock(redisError, nil, "")
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

func TestElongateURLRedisErrorSet(t *testing.T) {
	shortURL := "test"
	longURL := &models.URL{
		LongURL: "test-long",
		Hits:    5,
	}
	url, _ := longURL.ToJSON()
	redisError := fmt.Errorf("fail")
	redisMock := redis.CreateRedisMock(nil, redisError, url)
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

func TestGetURLHits(t *testing.T) {
	shortURL := "test"
	longURL := &models.URL{
		LongURL: "test-long",
		Hits:    1,
	}
	url, _ := longURL.ToJSON()
	redisError := fmt.Errorf("fail")
	redisMock := redis.CreateRedisMock(nil, redisError, url)
	url_repo := &repos.URLRepo{
		RedisConn: redisMock,
	}

	hits, err := GetURLHits(url_repo, shortURL)

	if err != nil {
		t.Log("expected no errors but got", err)
		t.Fail()
	}

	if hits != 1 {
		t.Log("expected hits to be 1 but got", hits)
		t.Fail()
	}
}

func TestGetURLHitsRedisErrorGet(t *testing.T) {
	shortURL := "test"
	redisError := fmt.Errorf("fail")
	redisMock := redis.CreateRedisMock(redisError, nil, "")
	url_repo := &repos.URLRepo{
		RedisConn: redisMock,
	}

	hits, err := GetURLHits(url_repo, shortURL)

	if err != redisError {
		t.Log("expected error to be 'fail' but got", err)
		t.Fail()
	}

	if hits != 0 {
		t.Log("expected hits to be 0 but got", hits)
		t.Fail()
	}
}
