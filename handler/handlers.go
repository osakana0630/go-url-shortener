package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/osakana0630/go-url-shortener/shortener"
	"github.com/osakana0630/go-url-shortener/store"
	"net/http"
)

type UrlCreateRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var createRequest UrlCreateRequest
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(createRequest.LongUrl, createRequest.UserId)
	store.SaveUrlMapping(shortUrl, createRequest.LongUrl, createRequest.UserId)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
