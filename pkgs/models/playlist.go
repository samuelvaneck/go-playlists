package model

import (
	"time"
)

type Playlist struct {
	ID             int       `json:"id"`
	SongID         int       `json:"song_id"`
	RadioStationID int       `json:"radio_station_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// func PlaylistById(id int) (Playlist, error) {
// 	var playlist Playlist
// 	var err error

// 	db, dbConnectionError := initializers.ConnectToDB()

// 	if dbConnectionError != nil {
// 		log.Fatalf("Unable to connection to database: %v", err)
// 		err = dbConnectionError
// 	}
// 	queryError := db.Find(&playlist, id).Error
// 	if queryError != nil {
// 		err = queryError
// 	}

// 	return playlist, err
// }

// func Playlists() ([]Playlist, error) {
// 	var playlists []Playlist
// 	var err error

// 	db, dbConnectionError := initializers.ConnectToDB()
// 	if dbConnectionError != nil {
// 		log.Fatalf("Unable to connection to database: %v", err)
// 		err = dbConnectionError
// 	}

// 	queryError := db.Find(&playlists).Error
// 	if queryError != nil {
// 		err = queryError
// 	}

// 	return playlists, err
// }
