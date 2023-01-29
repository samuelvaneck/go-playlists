package model

import (
	"context"
	"log"
	"playlists/pkgs/initializers"
	"time"
)

type RadioStation struct {
	Id                int64     `json:"id"`
	Name              string    `json:"name"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Api_url           string    `json:"api_url"`
	Stream_url        string    `json:"stream_url"`
	Scraper_processor string    `json:"screaper_processor"`
	Scraper_import    bool      `json:"scraper_import"`
}

func RadioStationById(id int) (RadioStation, error) {
	var radioStation RadioStation
	var queryErr error
	var err error

	conn, dbConnectionError := initializers.ConnectToDB()
	if dbConnectionError != nil {
		log.Fatalf("Unable to connection to database: %v", err)
		err = dbConnectionError
	}
	queryErr = conn.QueryRow(context.Background(), "SELECT * FROM radio_stations WHERE id = $1", id).Scan(&radioStation.Id, &radioStation.Name, &radioStation.CreatedAt, &radioStation.UpdatedAt, &radioStation.Api_url, &radioStation.Stream_url, &radioStation.Scraper_processor, &radioStation.Scraper_import)
	if queryErr != nil {
		log.Fatalf("Unable to query database: %v", err)
		err = queryErr
	}
	return radioStation, err
}

func RadioStations() ([]RadioStation, error) {
	var radioStations []RadioStation
	var err error

	conn, dbConnectionError := initializers.ConnectToDB()
	if dbConnectionError != nil {
		log.Fatalf("Unable to connection to database: %v", err)
		err = dbConnectionError
	}
	rows, queryError := conn.Query(context.Background(), "SELECT * FROM radio_stations")
	if err != nil {
		err = queryError
	}

	for rows.Next() {
		var radioStation RadioStation
		err := rows.Scan(&radioStation.Id, &radioStation.Name, &radioStation.CreatedAt, &radioStation.UpdatedAt, &radioStation.Api_url, &radioStation.Stream_url, &radioStation.Scraper_processor, &radioStation.Scraper_import)
		if err != nil {
			log.Fatal(err)
		}
		radioStations = append(radioStations, radioStation)
	}

	return radioStations, err
}
