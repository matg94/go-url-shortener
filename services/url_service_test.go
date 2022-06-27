package services

import (
	"fmt"
	"testing"

	"github.com/matg94/go-url-shortener/redis"
)

func TestShortenURL(t *testing.T) {
	longURL := "test"
	hashLength := 10
	redisConn := &redis.RedisConnectionMock{}
	redisConn.ReturnValue = ""

	shortened, err := ShortenURL(redisConn, longURL, hashLength)
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
	redisConn := &redis.RedisConnectionMock{}
	redisConn.ReturnValue = ""
	redisError := fmt.Errorf("fail")
	redisConn.ReturnErrorGet = redisError

	shortened, err := ShortenURL(redisConn, longURL, hashLength)
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
	redisConn := &redis.RedisConnectionMock{}
	redisConn.ReturnValue = ""
	redisError := fmt.Errorf("fail")
	redisConn.ReturnErrorSet = redisError

	shortened, err := ShortenURL(redisConn, longURL, hashLength)
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
	redisConn := &redis.RedisConnectionMock{}
	redisConn.ReturnValue = "test-long"

	url, err := ElongateURL(redisConn, shortURL)

	if err != nil {
		t.Log("expected error to be nil but got", err)
		t.Fail()
	}

	if url != "test-long" {
		t.Log("expected url to be 'test-long' but got", url)
		t.Fail()
	}
}

func TestElongateURLRedisError(t *testing.T) {
	shortURL := "test"
	redisConn := &redis.RedisConnectionMock{}
	redisConn.ReturnValue = "test-long"
	redisError := fmt.Errorf("fail")
	redisConn.ReturnErrorGet = redisError

	url, err := ElongateURL(redisConn, shortURL)

	if err != redisError {
		t.Log("expected error to be 'fail' but got", err)
		t.Fail()
	}

	if url != "" {
		t.Log("expected url to be empty but got", url)
		t.Fail()
	}
}
