package main

import (
	"fmt"
	"log"
	"os"
	"playlists/pkgs/handlers"
	"playlists/pkgs/initializers"
	models "playlists/pkgs/models"
	"playlists/pkgs/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.ConnectToRedis()
}

func main() {
	DB, err := initializers.ConnectToDB()
	redisConnection, err := initializers.ConnectToRedis()

	fmt.Fprint(os.Stdout, "Redis connection: ", redisConnection, "")

	if err != nil {
		log.Fatalf("Unable to connection to database: %v", err)
	}
	DB.AutoMigrate(&models.RadioStation{}, &models.Playlist{}, &models.Song{}, &models.Artist{})

	h := handlers.New(DB)
	app := fiber.New()
	routes.AllRoutes(app, h)

	fmt.Println("Listening on port", os.Getenv("APP_PORT"))
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
