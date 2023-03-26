package main

import (
	"fmt"
	"log"
	"os"
	"playlists/pkgs/handlers"
	"playlists/pkgs/initializers"
	models "playlists/pkgs/models"
	"playlists/pkgs/routes"
	services "playlists/pkgs/services"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var RedisClient *redis.Client
var DB *gorm.DB

func init() {
	var dbError error
	var redisError error

	initializers.LoadEnv()
	DB, dbError = initializers.ConnectToDB()
	if dbError != nil {
		log.Fatalf("Unable to connect to database: %v", dbError)
	}
	RedisClient, redisError = initializers.ConnectToRedis()
	if redisError != nil {
		log.Fatalf("Unable to connect to Redis: %v", redisError)
	}
}

func main() {
	DB.AutoMigrate(&models.RadioStation{}, &models.Playlist{}, &models.Song{}, &models.Artist{})

	h := handlers.New(DB)
	r := models.NewRedisClient(RedisClient)
	r.Client.Ping(r.Ctx) // temp
	app := fiber.New()
	routes.AllRoutes(app, h)

	go services.SongRecognizer()

	fmt.Println("Listening on port", os.Getenv("APP_PORT"))
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
