package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matg94/go-url-shortener/controllers"
)

func initializeRoutes(serv *gin.Engine) {
	serv.POST("/shorten", controllers.PostShortenURL)
	serv.POST("/elong", controllers.PostLongURL)
	serv.GET("/hits/:hash", controllers.GetURLHits)
	serv.GET("/:hash", controllers.ShortRedirect)
}
