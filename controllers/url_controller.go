package controllers

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/matg94/go-url-shortener/config"
	"github.com/matg94/go-url-shortener/models"
	"github.com/matg94/go-url-shortener/repos"
	"github.com/matg94/go-url-shortener/services"
)

var URLRepo repos.URLRepoInterface
var AppConfig config.AppConfig

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
	request := models.ShortenRequestFromJson(requestBody)
	shortened, err := services.ShortenURL(URLRepo, request.URL, AppConfig.HashLength)
	if err != nil {
		HandleError(c, err, 500)
	}
	c.JSON(200, gin.H{
		"URL": shortened,
	})
}

func GetLongURL(c *gin.Context) {
	body := c.Request.Body
	requestBody, readError := ioutil.ReadAll(body)
	if readError != nil {
		HandleError(c, readError, 400)
	}
	request := models.LongRequestFromJson(requestBody)
	longURL, err := services.ElongateURL(URLRepo, request.URL)
	if err != nil {
		HandleError(c, err, 500)
	}
	c.JSON(200, gin.H{
		"URL": longURL,
	})
}

func GetURLHits(c *gin.Context) {

}

func ShortRedirect(c *gin.Context) {
	shortenedURL := c.Param("url")
	originalURL, err := services.ElongateURL(URLRepo, shortenedURL)
	if err != nil {
		HandleError(c, err, 400)
	}
	c.Redirect(302, originalURL)
}

/*
	POST: /shorten -> {}


*/
