package main

import (
	"fmt"
	"log"
	"os"
	"playlists/pkgs/handlers"
	"playlists/pkgs/initializers"
	models "playlists/pkgs/models"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	DB, err := initializers.ConnectToDB()
	if err != nil {
		log.Fatalf("Unable to connection to database: %v", err)
	}
	DB.AutoMigrate(&models.RadioStation{}, &models.Playlist{}, &models.Song{})

	h := handlers.New(DB)
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/radio_stations", h.GetAllRadioStations)
	app.Get("/api/radio_stations/:id", h.GetRadioStation)
	app.Get("/api/playlists", h.GetAllPlaylists)
	app.Get("/api/playlists/:id", h.GetPlaylist)
	app.Get("/api/songs", h.GetAllSongs)
	app.Get("/api/songs/:id", h.GetSong)

	fmt.Println("Listening on port", os.Getenv("APP_PORT"))
}
