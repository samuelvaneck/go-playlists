package models

import "gorm.io/gorm"

type Song struct {
	ID                  int64          `json:"id"`
	Title               string         `json:"title"`
	Isrc                string         `json:"isrc"`
	ID_on_spotify       string         `json:"id_on_spotify"`
	Spotify_song_url    string         `json:"spotify_song_url"`
	Spotify_artwork_url string         `json:"spotify_artwork_url"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
