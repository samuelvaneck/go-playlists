package services

import (
	"encoding/json"
	"fmt"
	"log"
	models "playlists/pkgs/models"
	"time"
)

type Data struct {
	Track struct {
		Title    string `json:"title"`
		SubTitle string `json:"subtitle"`
	} `json:"track"`
}

func SongRecognizer() {
	var err error
	var radioStations []models.RadioStation
	var d Data

	ticker := time.NewTicker(1 * time.Minute)
	done := make(chan bool)

	radioStations, err = models.AllRadioStations()
	if err != nil {
		fmt.Println("Unable to get radio stations")
	}

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				for _, radioStation := range radioStations {
					recognizedSong, recErr := radioStation.RecognizeSong()
					if recErr != nil {
						log.Fatalf("Unable to recognize song: %v", recErr)
						err = recErr
					}

					parseErr := json.Unmarshal([]byte(recognizedSong), &d)
					if parseErr != nil {
						log.Fatalf("Unable to parse json: %v", parseErr)
						err = parseErr
					}

					fmt.Println(" ")
					fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
					fmt.Printf("Song playing on: %v\n", radioStation.Name)
					fmt.Println(d.Track.Title, " from ", d.Track.SubTitle)
					fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
				}
				fmt.Println("All done, waiting 1 minute...")
				time.Sleep(1 * time.Minute)
			}
		}
	}()

	time.Sleep(5 * time.Minute)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
