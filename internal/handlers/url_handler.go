package handlers

import (
	"net/http"
	"os"

	"urlshortener/internal/database"
	"urlshortener/internal/models"
	"urlshortener/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateShortURL(c *gin.Context) {
	var request models.CreateURLRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortCode, err := utils.GenerateShortCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short code"})
		return
	}

	for {
		var existingURL models.URL
		if err := database.DB.Where("short_code = ?", shortCode).First(&existingURL).Error; err != nil {
			break
		}
		shortCode, err = utils.GenerateShortCode()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short code"})
			return
		}
	}

	url := models.URL{
		OriginalURL: request.URL,
		ShortCode:   shortCode,
		Clicks:      0,
	}

	if err := database.DB.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create short URL"})
		return
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	response := models.URLResponse{
		ID:          url.ID,
		OriginalURL: url.OriginalURL,
		ShortCode:   url.ShortCode,
		ShortURL:    baseURL + "/" + url.ShortCode,
		Clicks:      url.Clicks,
		CreatedAt:   url.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func RedirectURL(c *gin.Context) {
	shortCode := c.Param("code")

	var url models.URL
	if err := database.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	database.DB.Model(&url).Update("clicks", url.Clicks+1)

	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}

func GetURLStats(c *gin.Context) {
	shortCode := c.Param("code")

	var url models.URL
	if err := database.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	response := models.URLResponse{
		ID:          url.ID,
		OriginalURL: url.OriginalURL,
		ShortCode:   url.ShortCode,
		ShortURL:    baseURL + "/" + url.ShortCode,
		Clicks:      url.Clicks,
		CreatedAt:   url.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func GetAllURLs(c *gin.Context) {
	var urls []models.URL
	if err := database.DB.Find(&urls).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch URLs"})
		return
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	var responses []models.URLResponse
	for _, url := range urls {
		response := models.URLResponse{
			ID:          url.ID,
			OriginalURL: url.OriginalURL,
			ShortCode:   url.ShortCode,
			ShortURL:    baseURL + "/" + url.ShortCode,
			Clicks:      url.Clicks,
			CreatedAt:   url.CreatedAt,
		}
		responses = append(responses, response)
	}

	c.JSON(http.StatusOK, responses)
}
