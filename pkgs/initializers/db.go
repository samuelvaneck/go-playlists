package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() (*gorm.DB, error) {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return nil, err
	}

	log.Println("Connected to database")

	return DB, nil
}
