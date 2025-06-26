package main

import (
	"log"
	"os"

	"urlshortener/internal/database"
	"urlshortener/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("config.env"); err != nil {
		log.Println("No config.env file found, using environment variables")
	}

	database.Connect()

	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	r.GET("/styles.css", func(c *gin.Context) {
		c.File("./static/styles.css")
	})
	r.GET("/script.js", func(c *gin.Context) {
		c.File("./static/script.js")
	})

	api := r.Group("/api/v1")
	{
		api.POST("/shorten", handlers.CreateShortURL)
		api.GET("/urls", handlers.GetAllURLs)
		api.GET("/stats/:code", handlers.GetURLStats)
	}

	// Redirect handler phải đặt cuối cùng
	r.GET("/:code", handlers.RedirectURL)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}
