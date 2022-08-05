package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/matg94/go-url-shortener/config"
	"github.com/matg94/go-url-shortener/controllers"
	"github.com/matg94/go-url-shortener/redis"
	"github.com/matg94/go-url-shortener/repos"
)

func LoadConfig(profile string) *config.AppConfig {
	data, err := config.ReadConfigFile(fmt.Sprintf("./config/%s.yaml", profile))
	if err != nil {
		log.Fatalf("Could not read config file for profile: %s", profile)
	}
	appConfig, err := config.ParseYamlConfig(data)
	if err != nil {
		log.Fatalf("Could not parse config for profile: %s", profile)
	}
	return appConfig
}

func setupDataPersistence(redisConfig config.RedisConfig) redis.RedisConnectionInterface {
	if redisConfig.IsCache {
		log.Print("Starting data persistence with CACHE")
		return redis.CreateRedisCache()
	}
	log.Printf("Starting data persistence with redis at %s:%d", redisConfig.URL, redisConfig.Port)
	return redis.CreateRedisConnectionPool(&redisConfig)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Application profile is required as an Arg to run")
	}
	profile := os.Args[1]
	appConfig := LoadConfig(profile)
	controllers.AppConfig = *appConfig
	dataPersistence := setupDataPersistence(appConfig.Redis)
	controllers.URLRepo = repos.CreateURLRepo(dataPersistence)

	server := gin.Default()
	server.Use(cors.Default())
	initializeRoutes(server)
	err := server.Run()
	if err != nil {
		log.Fatalf("Failed to run server! %s", err)
	}
}
