package models

import "gorm.io/gorm"

type Playlist struct {
	ID             int64          `json:"id"`
	SongID         int64          `json:"song_id"`
	RadioStationID int64          `json:"radio_station_id"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Song           Song           `json:"song"`
}
