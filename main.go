package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/matg94/go-url-shortener/config"
	"github.com/matg94/go-url-shortener/controllers"
	"github.com/matg94/go-url-shortener/models"
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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Application profile is required as an Arg to run")
	}
	profile := os.Args[1]
	appConfig := LoadConfig(profile)
	controllers.AppConfig = *appConfig
	redisMock := &redis.RedisConnectionMock{}
	returnedUrl := models.URL{
		LongURL: "Hello!",
		Hits:    3,
	}
	returnedString, _ := returnedUrl.ToJSON()
	redisMock.ReturnValue = returnedString
	controllers.URLRepo = repos.CreateURLRepo(redisMock)

	server := gin.Default()
	server.Use(cors.Default())
	initializeRoutes(server)
	server.Run()
}
