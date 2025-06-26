package models

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	OriginalURL string         `gorm:"not null" json:"original_url" binding:"required"`
	ShortCode   string         `gorm:"uniqueIndex;not null" json:"short_code"`
	Clicks      int64          `gorm:"default:0" json:"clicks"`
}

type CreateURLRequest struct {
	URL string `json:"url" binding:"required,url"`
}

type URLResponse struct {
	ID          uint      `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`
	ShortURL    string    `json:"short_url"`
	Clicks      int64     `json:"clicks"`
	CreatedAt   time.Time `json:"created_at"`
}
