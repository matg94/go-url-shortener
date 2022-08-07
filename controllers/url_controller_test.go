package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/matg94/go-url-shortener/config"
	"github.com/matg94/go-url-shortener/models"
	"github.com/matg94/go-url-shortener/redis"
	"github.com/matg94/go-url-shortener/repos"
)

func SetupRepo() {
	appConfig := config.LoadConfig("test")
	AppConfig = *appConfig
	redisCache := redis.SetupDataPersistence(appConfig.Redis)
	URLRepo = repos.CreateURLRepo(redisCache)
}

func AddRedisTestData() {
	err := URLRepo.StoreURL("234988566c9a0a9cf952", models.URL{
		LongURL: "http://google.com",
		Hits:    3,
	})
	fmt.Println(err)
}

// POST SHORTEN URL

func TestPostShortenURL(t *testing.T) {
	SetupRepo()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := strings.NewReader("{\"url\": \"http://google.com\"}")
	c.Request, _ = http.NewRequest("POST", "http://localhost:8080", body)

	PostShortenURL(c)

	if w.Code != 200 {
		t.Log("expected 200 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != "{\"URL\":\"234988566c9a0a9cf952\"}" {
		t.Log("expected url 234988566c9a0a9cf952 but got response", w.Body.String())
		t.Fail()
	}
}

func TestPostShortenURLMissingURL(t *testing.T) {
	SetupRepo()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := strings.NewReader("{\"abc\": \"http://google.com\"}")
	c.Request, _ = http.NewRequest("POST", "http://localhost:8080", body)

	PostShortenURL(c)

	if w.Code != 400 {
		t.Log("expected 400 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != "{\"error\":\"URL not defined in request\",\"status\":400}" {
		t.Log("expected url URL not defined in request but got response", w.Body.String())
		t.Fail()
	}
}

func TestPostShortenURLInvalidRequest(t *testing.T) {
	SetupRepo()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := strings.NewReader("s{\"}")
	c.Request, _ = http.NewRequest("POST", "http://localhost:8080", body)

	PostShortenURL(c)

	if w.Code != 400 {
		t.Log("expected 400 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != "{\"error\":\"invalid character 's' looking for beginning of value\",\"status\":400}" {
		t.Log("expected url URL not defined in request but got response", w.Body.String())
		t.Fail()
	}
}

func TestPostShortenURLInvalidConfig(t *testing.T) {
	SetupRepo()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	redisMock := redis.CreateRedisMock(errors.New("some redis error"), nil, "some value")
	URLRepo = repos.CreateURLRepo(redisMock)

	body := strings.NewReader("{\"url\": \"http://google.com\"}")
	c.Request, _ = http.NewRequest("POST", "http://localhost:8080", body)

	PostShortenURL(c)

	if w.Code != 500 {
		t.Log("expected 500 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != "{\"error\":\"some redis error\",\"status\":500}" {
		t.Log("expected url some redis error but got response", w.Body.String())
		t.Fail()
	}
}

// POST LONG URL

func TestPostLongURL(t *testing.T) {
	SetupRepo()
	AddRedisTestData()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := strings.NewReader("{\"url\": \"234988566c9a0a9cf952\"}")
	c.Request, _ = http.NewRequest("POST", "http://localhost:8080", body)

	PostLongURL(c)

	if w.Code != 200 {
		t.Log("expected 200 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != "{\"URL\":\"http://google.com\"}" {
		t.Log("expected url http://google.com but got response", w.Body.String())
		t.Fail()
	}
}

func TestPostLongURLNotFound(t *testing.T) {
	SetupRepo()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := strings.NewReader("{\"url\": \"234988566c9a0a9cf952\"}")
	c.Request, _ = http.NewRequest("POST", "http://localhost:8080", body)

	PostLongURL(c)

	if w.Code != 404 {
		t.Log("expected 404 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != "{\"error\":\"short url key not found\",\"status\":404}" {
		t.Log("expected url not found but got response", w.Body.String())
		t.Fail()
	}
}

func TestPostLongURLMissingURL(t *testing.T) {
	SetupRepo()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := strings.NewReader("{\"abc\": \"http://google.com\"}")
	c.Request, _ = http.NewRequest("POST", "http://localhost:8080", body)

	PostLongURL(c)

	if w.Code != 400 {
		t.Log("expected 400 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != "{\"error\":\"URL not defined in request\",\"status\":400}" {
		t.Log("expected url URL not defined in request but got response", w.Body.String())
		t.Fail()
	}
}

func TestPostLongURLRedisError(t *testing.T) {
	SetupRepo()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := strings.NewReader("{\"url\": \"234988566c9a0a9cf952\"}")
	c.Request, _ = http.NewRequest("POST", "http://localhost:8080", body)

	redisMock := redis.CreateRedisMock(errors.New("some redis error"), nil, "some value")
	URLRepo = repos.CreateURLRepo(redisMock)

	PostLongURL(c)

	if w.Code != 500 {
		t.Log("expected 500 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != "{\"error\":\"some redis error\",\"status\":500}" {
		t.Log("expected url some redis error but got response", w.Body.String())
		t.Fail()
	}
}

// GET URL HITS

func TestGetURLHits(t *testing.T) {
	SetupRepo()
	AddRedisTestData()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "http://localhost/hits/234988566c9a0a9cf952", nil)
	c.Params = gin.Params{gin.Param{
		Key:   "hash",
		Value: "234988566c9a0a9cf952",
	}}

	GetURLHits(c)

	if w.Code != 200 {
		t.Log("expected 200 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != "{\"Hits\":3,\"URL\":\"234988566c9a0a9cf952\"}" {
		t.Log("expected Hits 3 & url 234988566c9a0a9cf952 but got response", w.Body.String())
		t.Fail()
	}
}

func TestGetURLHitsURLNotFound(t *testing.T) {
	SetupRepo()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "http://localhost/hits/234988566c9a0a9cf952", nil)
	c.Params = gin.Params{gin.Param{
		Key:   "hash",
		Value: "234988566c9a0a9cf952",
	}}

	GetURLHits(c)

	if w.Code != 404 {
		t.Log("expected 404 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != "{\"error\":\"short url key not found\",\"status\":404}" {
		t.Log("expected not found error but got response", w.Body.String())
		t.Fail()
	}
}

func TestGetURLHitsRedisError(t *testing.T) {
	SetupRepo()
	AddRedisTestData()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "http://localhost/hits/234988566c9a0a9cf952", nil)
	c.Params = gin.Params{gin.Param{
		Key:   "hash",
		Value: "234988566c9a0a9cf952",
	}}

	redisMock := redis.CreateRedisMock(errors.New("some redis error"), nil, "some value")
	URLRepo = repos.CreateURLRepo(redisMock)

	GetURLHits(c)

	if w.Code != 500 {
		t.Log("expected 500 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != "{\"error\":\"some redis error\",\"status\":500}" {
		t.Log("expected url some redis error but got response", w.Body.String())
		t.Fail()
	}
}

func TestShortURLRedirect(t *testing.T) {
	SetupRepo()
	AddRedisTestData()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "http://localhost/234988566c9a0a9cf952", nil)
	c.Params = gin.Params{gin.Param{
		Key:   "hash",
		Value: "234988566c9a0a9cf952",
	}}

	ShortRedirect(c)

	if w.Code != 302 {
		t.Log("expected 302 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != "<a href=\"http://google.com\">Found</a>.\n\n" {
		t.Log("expected google href redirect but got", w.Body.String())
		t.Fail()
	}
}

func TestShortURLRedirectRedisError(t *testing.T) {
	SetupRepo()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "http://localhost/234988566c9a0a9cf952", nil)
	c.Params = gin.Params{gin.Param{
		Key:   "hash",
		Value: "234988566c9a0a9cf952",
	}}

	redisMock := redis.CreateRedisMock(errors.New("some redis error"), nil, "some value")
	URLRepo = repos.CreateURLRepo(redisMock)

	ShortRedirect(c)

	if w.Code != 500 {
		t.Log("expected 500 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != "{\"error\":\"some redis error\",\"status\":500}" {
		t.Log("expected url some redis error but got response", w.Body.String())
		t.Fail()
	}
}

func TestShortURLRedirectURLNotFound(t *testing.T) {
	SetupRepo()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "http://localhost/123", nil)
	c.Params = gin.Params{gin.Param{
		Key:   "hash",
		Value: "123",
	}}

	ShortRedirect(c)

	if w.Code != 404 {
		t.Log("expected 404 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != "{\"error\":\"short url key not found\",\"status\":404}" {
		t.Log("expected url some redis error but got response", w.Body.String())
		t.Fail()
	}
}
