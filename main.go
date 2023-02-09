package main

import (
	"fmt"
	"log"
	"os"
	handlers "playlists/pkgs/handlers"
	"playlists/pkgs/initializers"
	model "playlists/pkgs/models"

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
	DB.AutoMigrate(&model.RadioStation{}, &model.Playlist{})

	h := handlers.New(DB)
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/radio_stations", h.GetAllRadioStations)
	app.Get("/api/radio_stations/:id", h.GetRadioStation)
	app.Get("/api/playlists", h.GetAllPlaylists)
	app.Get("/api/playlists/:id", h.GetPlaylist)

	fmt.Println("Listening on port", os.Getenv("APP_PORT"))
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
