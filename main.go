package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohammadjafarsafarali/golang-link-shortener/handler"
	"github.com/mohammadjafarsafarali/golang-link-shortener/store"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("There is a problem with env file!")
	}
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	router.POST("/url-shortener", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	router.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	httpPort := os.Getenv("APP_PORT")
	httpHost := os.Getenv("APP_HOST")
	if httpPort == "" {
		httpPort = "8080"
	}
	if httpHost == "" {
		httpPort = "localhost"
	}

	err := router.Run(":" + httpPort)
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
