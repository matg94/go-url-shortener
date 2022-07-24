package controllers

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/matg94/go-url-shortener/config"
	"github.com/matg94/go-url-shortener/models"
	"github.com/matg94/go-url-shortener/repos"
	"github.com/matg94/go-url-shortener/services"
)

var url_repo *repos.URLRepo
var appConfig *config.AppConfig

func HandleError(c *gin.Context, err error, code int) {
	c.JSON(code, gin.H{
		"error": err.Error(),
	})
}

func PostShortenURL(c *gin.Context) {
	body := c.Request.Body
	requestBody, readError := ioutil.ReadAll(body)
	if readError != nil {
		HandleError(c, readError, 400)
	}
	request := models.ShortenRequestFromJson(string(requestBody))
	shortened, err := services.ShortenURL(url_repo, request.URL, appConfig.HashLength)
	if err != nil {
		HandleError(c, err, 500)
	}
	c.JSON(200, gin.H{
		"URL": shortened,
	})
}

func GetLongURL(c *gin.Context) {

}

func GetURLHits(c *gin.Context) {

}

/*
	POST: /shorten -> {}


*/
