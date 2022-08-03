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
		"status": code,
		"error":  err.Error(),
	})
}

func PostShortenURL(c *gin.Context) {
	body := c.Request.Body
	requestBody, readError := ioutil.ReadAll(body)
	if readError != nil {
		HandleError(c, readError, 400)
		return
	}
	request, parseError := models.ShortenRequestFromJson(requestBody)
	if parseError != nil {
		HandleError(c, parseError, 400)
		return
	}
	shortened, err := services.ShortenURL(URLRepo, request.URL, AppConfig.HashLength)
	if err != nil {
		HandleError(c, err, 500)
		return
	}
	c.JSON(200, gin.H{
		"URL": shortened,
	})
}

func PostLongURL(c *gin.Context) {
	body := c.Request.Body
	requestBody, readError := ioutil.ReadAll(body)
	if readError != nil {
		HandleError(c, readError, 400)
		return
	}
	request, parseError := models.LongRequestFromJson(requestBody)
	if parseError != nil {
		HandleError(c, parseError, 400)
		return
	}
	longURL, err := services.ElongateURL(URLRepo, request.Hash)
	if err != nil {
		HandleError(c, err, 500)
		return
	}
	c.JSON(200, gin.H{
		"URL": longURL,
	})
}

func GetURLHits(c *gin.Context) {
	hash := c.Param("hash")
	request := models.URLElongateResponse{
		Hash: hash,
	}
	longURL, err := services.ElongateURL(URLRepo, request.Hash)
	if err != nil {
		HandleError(c, err, 500)
		return
	}
	hits, err := services.GetURLHits(URLRepo, request.Hash)
	if err != nil {
		HandleError(c, err, 500)
		return
	}
	c.JSON(200, gin.H{
		"URL":  longURL,
		"Hits": hits,
	})
}

func ShortRedirect(c *gin.Context) {
	shortenedURL := c.Param("hash")
	originalURL, err := services.ElongateURL(URLRepo, shortenedURL)
	if err != nil {
		HandleError(c, err, 400)
		return
	}
	c.Redirect(302, originalURL)
}
