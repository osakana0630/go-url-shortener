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

// CreateShortUrl
// 元URLを受け取り、短縮URLを生成し、redisに保存し、短縮URLを返す
func CreateShortUrl(c *gin.Context) {
	var createRequest UrlCreateRequest
	// NOTE: ShouldBindJSON関数を使用して、リクエストボディをUrlCreateRequest構造体にバインドしている
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(createRequest.LongUrl, createRequest.UserId)
	store.SaveUrlMapping(shortUrl, createRequest.LongUrl, createRequest.UserId)

	host := "http://localhost:9808/"
	c.JSON(http.StatusOK, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

// HandleShortUrlRedirect
// 短縮URLを受け取り、元のURLにリダイレクトする
func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl") // pathパラメータから短縮URLを取得
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(http.StatusFound, initialUrl)
}
