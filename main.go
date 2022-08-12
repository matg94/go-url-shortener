package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/matg94/go-url-shortener/config"
	"github.com/matg94/go-url-shortener/controllers"
	"github.com/matg94/go-url-shortener/redis"
	"github.com/matg94/go-url-shortener/repos"
)

func main() {
	profile := os.Getenv("APP_CONFIG_PROFILE")
	appConfig := config.LoadConfig(profile)
	controllers.AppConfig = *appConfig
	dataPersistence := redis.SetupDataPersistence(appConfig.Redis)
	controllers.URLRepo = repos.CreateURLRepo(dataPersistence)

	server := gin.Default()
	server.Use(cors.Default())
	initializeRoutes(server)
	err := server.Run()
	if err != nil {
		log.Fatalf("Failed to run server! %s", err)
	}
}
