package model

import (
	"time"

	"gorm.io/gorm"
)

type RadioStation struct {
	ID                int64          `json:"id"`
	Name              string         `json:"name"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	Api_url           string         `json:"api_url"`
	Stream_url        string         `json:"stream_url"`
	Scraper_processor string         `json:"screaper_processor"`
	Scraper_import    bool           `json:"scraper_import"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
