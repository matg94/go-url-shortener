package redis

import (
	"testing"
)

func TestCacheGetSet(t *testing.T) {
	redisCache := CreateRedisCache()
	testVal := "test123"
	err := redisCache.SET("test", testVal)
	if err != nil {
		t.Logf("expected no errors but got %s", err)
		t.Fail()
	}
	val, err := redisCache.GET("test")
	if err != nil {
		t.Logf("expected no errors but got %s", err)
		t.Fail()
	}
	if val != testVal {
		t.Logf("expected GET to return %s but got %s", testVal, val)
		t.Fail()
	}
}
