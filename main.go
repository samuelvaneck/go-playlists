package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"playlists/pkgs/controllers"
	"playlists/pkgs/initializers"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	conn, err := initializers.ConnectToDB()
	if err != nil {
		log.Fatalf("Unable to connection to database: %v", err)
	}

	pingErr := conn.Ping(context.Background())
	if pingErr != nil {
		log.Fatalf("Unable to ping database: %v", pingErr)
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/radio_stations", controllers.RadioStationsIndex)
	app.Get("/api/radio_stations/:id", controllers.RadioStationById)

	fmt.Println("Listening on port", os.Getenv("APP_PORT"))
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
