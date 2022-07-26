package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matg94/go-url-shortener/controllers"
)

func initializeRoutes(serv *gin.Engine) {
	serv.POST("/url", controllers.PostShortenURL)
	serv.GET("/url", controllers.GetLongURL)
	serv.GET("/hits", controllers.GetURLHits)
	serv.GET("/:url", controllers.ShortRedirect)
}
