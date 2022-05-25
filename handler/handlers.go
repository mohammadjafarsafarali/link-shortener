package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammadjafarsafarali/golang-link-shortener/shortener"
	"github.com/mohammadjafarsafarali/golang-link-shortener/store"
	"net/http"
	"os"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl)

	//store
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	host := os.Getenv("SHORT_LINK_HOST")
	c.JSON(http.StatusOK, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	originalUrl := store.RetrieveOriginalUrl(shortUrl)
	c.Redirect(302, originalUrl)
}
