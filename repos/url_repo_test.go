package repos

import (
	"fmt"
	"testing"

	"github.com/matg94/go-url-shortener/models"
	"github.com/matg94/go-url-shortener/redis"
)

func TestStoreURL(t *testing.T) {
	longURL := "test-long"
	shortURL := "test-short"
	hits := uint(1)

	redisMock := &redis.RedisConnectionMock{}

	repo := &URLRepo{
		RedisConn: redisMock,
	}

	err := repo.StoreURL(shortURL, models.URL{
		LongURL: longURL,
		Hits:    hits,
	})
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

func TestStoreURLRedisErrorSet(t *testing.T) {
	longURL := "test-long"
	shortURL := "test-short"
	hits := uint(1)

	redisMock := &redis.RedisConnectionMock{}
	redisError := fmt.Errorf("fail")
	redisMock.ReturnErrorSet = redisError
	repo := &URLRepo{
		RedisConn: redisMock,
	}

	err := repo.StoreURL(shortURL, models.URL{
		LongURL: longURL,
		Hits:    hits,
	})
	if err != redisError {
		t.Logf("expected error %s but got %s", redisError, err)
		t.Fail()
	}
}

func TestGetURL(t *testing.T) {
	longURL := "test-long"
	shortURL := "test-short"
	hits := uint(1)

	redisMock := &redis.RedisConnectionMock{}
	returnedUrl := models.URL{
		LongURL: longURL,
		Hits:    hits,
	}

	repo := &URLRepo{
		RedisConn: redisMock,
	}
	returnedString, _ := returnedUrl.ToJSON()
	redisMock.ReturnValue = returnedString

	val, err := repo.GetURL(shortURL)
	if err != nil {
		t.Logf("expected no errors but got %s", err)
		t.Fail()
	}
	valString, err := val.ToJSON()
	if err != nil {
		t.Logf("expected returned value to be serializable but got error %s", err)
		t.Fail()
	}

	if valString != returnedString {
		t.Logf("expected returned value to match %s, but got %s", returnedString, valString)
		t.Fail()
	}
}

func TestGetURLRedisErrorGet(t *testing.T) {
	longURL := "test-long"
	shortURL := "test-short"
	hits := uint(1)

	redisMock := &redis.RedisConnectionMock{}
	redisError := fmt.Errorf("fail")
	redisMock.ReturnErrorGet = redisError
	returnedUrl := models.URL{
		LongURL: longURL,
		Hits:    hits,
	}

	repo := &URLRepo{
		RedisConn: redisMock,
	}
	returnedString, _ := returnedUrl.ToJSON()
	redisMock.ReturnValue = returnedString

	returnedVal, err := repo.GetURL(shortURL)
	if err != redisError {
		t.Logf("expected fail error but got %s", err)
		t.Fail()
	}

	emptyURL := models.URL{}
	returnedValString, _ := returnedVal.ToJSON()
	if returnedVal != emptyURL {
		t.Logf("expected empty URL returned but got %s", returnedValString)
		t.Fail()
	}
}

func TestGetURLNoResult(t *testing.T) {
	longURL := "test-long"
	shortURL := "test-short"
	hits := uint(1)

	redisMock := &redis.RedisConnectionMock{}
	redisError := redis.ErrRedisValueNotFound
	redisMock.ReturnErrorGet = redisError
	returnedUrl := models.URL{
		LongURL: longURL,
		Hits:    hits,
	}

	repo := &URLRepo{
		RedisConn: redisMock,
	}
	returnedString, _ := returnedUrl.ToJSON()
	redisMock.ReturnValue = returnedString

	returnedVal, err := repo.GetURL(shortURL)
	if err != redisError {
		t.Logf("expected not found error but got %s", err)
		t.Fail()
	}

	emptyURL := models.URL{}
	returnedValString, _ := returnedVal.ToJSON()
	if returnedVal != emptyURL {
		t.Logf("expected empty URL returned but got %s", returnedValString)
		t.Fail()
	}
}

func TestIncrementHits(t *testing.T) {
	longURL := "test-long"
	shortURL := "test-short"
	hits := uint(5)

	redisMock := &redis.RedisConnectionMock{}
	returnedUrl := models.URL{
		LongURL: longURL,
		Hits:    hits,
	}
	returnedString, _ := returnedUrl.ToJSON()
	redisMock.ReturnValue = returnedString

	repo := &URLRepo{
		RedisConn: redisMock,
	}

	err := repo.IncrementHits(shortURL)
	if err != nil {
		t.Logf("expected no errors but got %s", err)
		t.Fail()
	}

	returnedURL, err := models.FromJSON(redisMock.Value)
	if err != nil {
		t.Logf("failed to parse json data %s", err)
		t.Fail()
	}

	if returnedURL.Hits != hits+1 {
		t.Logf("expected hits to be %d, but got %d", hits+1, returnedURL.Hits)
		t.Fail()
	}
}

func TestIncrementHitsGetRedisErr(t *testing.T) {
	longURL := "test-long"
	shortURL := "test-short"
	hits := uint(5)

	redisMock := &redis.RedisConnectionMock{}
	returnedUrl := models.URL{
		LongURL: longURL,
		Hits:    hits,
	}
	returnedString, _ := returnedUrl.ToJSON()
	redisMock.ReturnValue = returnedString
	redisError := redis.ErrRedisValueNotFound
	redisMock.ReturnErrorGet = redisError

	repo := &URLRepo{
		RedisConn: redisMock,
	}

	err := repo.IncrementHits(shortURL)

	if err != redisError {
		t.Logf("expected not found error but got %s", err)
		t.Fail()
	}
}

func TestIncrementHitsSetRedisErr(t *testing.T) {
	longURL := "test-long"
	shortURL := "test-short"
	hits := uint(5)

	redisMock := &redis.RedisConnectionMock{}
	returnedUrl := models.URL{
		LongURL: longURL,
		Hits:    hits,
	}
	returnedString, _ := returnedUrl.ToJSON()
	redisMock.ReturnValue = returnedString
	redisError := fmt.Errorf("fail")
	redisMock.ReturnErrorSet = redisError

	repo := &URLRepo{
		RedisConn: redisMock,
	}

	err := repo.IncrementHits(shortURL)

	if err != redisError {
		t.Logf("expected fail to set error but got %s", err)
		t.Fail()
	}
}
