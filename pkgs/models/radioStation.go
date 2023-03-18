package models

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"playlists/pkgs/initializers"

	"gorm.io/gorm"
)

type RadioStation struct {
	ID               int64          `json:"id"`
	Name             string         `json:"name"`
	ApiUrl           string         `json:"api_url"`
	StreamUrl        string         `json:"stream_url"`
	ScraperProcessor string         `json:"screaper_processor"`
	ScraperImport    bool           `json:"scraper_import"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (r RadioStation) RecognizeSong() (string, error) {
	var err error
	outputFile := r.outputFile()

	streamErr := r.CaptureStream()
	if streamErr != nil {
		log.Fatalf("Unable to capture stream: %v", err)
		return "", streamErr
	}

	out, songRecErr := exec.Command("songrec", "audio-file-to-recognized-song", outputFile).Output()
	if songRecErr != nil {
		log.Fatal(songRecErr)
		return "", songRecErr
	}

	return string(out), nil
}

func (r *RadioStation) CaptureStream() error {
	outputFile := r.outputFile()
	cmd := exec.Command("ffmpeg", "-y", "-t", "00:00:05", "-i", r.StreamUrl, "-c", "copy", outputFile)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (r RadioStation) outputFile() string {
	return fmt.Sprintf("tmp/%s.mp3", r.Name)
}

func GetRadioStation(id int64) (RadioStation, error) {
	var radioStation RadioStation
	db, err := initializers.ConnectToDB()
	if err != nil {
		return radioStation, err
	}

	queryErr := db.First(&radioStation, id).Error
	if queryErr != nil {
		return radioStation, err
	}
	return radioStation, nil
}

func AllRadioStations() ([]RadioStation, error) {
	var radioStations []RadioStation
	db, err := initializers.ConnectToDB()
	if err != nil {
		return nil, err
	}

	queryErr := db.Find(&radioStations).Error
	if queryErr != nil {
		return nil, err
	}
	return radioStations, nil
}
