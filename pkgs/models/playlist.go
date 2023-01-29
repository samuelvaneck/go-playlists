package model

import (
	"context"
	"log"
	"playlists/pkgs/initializers"
	"time"
)

type Playlist struct {
	ID             int64     `json:"id"`
	songID         int64     `json:"song_id"`
	radioStationID int64     `json:"radio_station_id"`
	createdAt      time.Time `json:"created_at"`
	updatedAt      time.Time `json:"updated_at"`
}

func PlaylistById(id int) (Playlist, error) {
	var playlist Playlist
	var err error

	conn, dbConnectionError := initializers.ConnectToDB()
	if dbConnectionError != nil {
		log.Fatalf("Unable to connection to database: %v", err)
		err = dbConnectionError
	}
	row, queryError := conn.Query(context.Background(), "SELECT * FROM playlists WHERE id = $1", id)
	row.Scan(&playlist.ID, &playlist.songID, &playlist.radioStationID, &playlist.createdAt, &playlist.updatedAt)
	if err != nil {
		err = queryError
	}
	return playlist, err
}

func Playlists() ([]Playlist, error) {
	var playlists []Playlist
	var err error

	conn, dbConnectionError := initializers.ConnectToDB()
	if dbConnectionError != nil {
		log.Fatalf("Unable to connection to database: %v", err)
		err = dbConnectionError
	}
	rows, queryError := conn.Query(context.Background(), "SELECT * FROM playlists")
	if err != nil {
		err = queryError
	}

	for rows.Next() {
		var playlist Playlist
		err = rows.Scan(&playlist.ID, &playlist.songID, &playlist.radioStationID, &playlist.createdAt, &playlist.updatedAt)
		if err != nil {
			log.Fatal(err)
		}
		playlists = append(playlists, playlist)
	}

	return playlists, err
}
