package initializers

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectToDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connection to database: %v", err)
	}
	return conn, nil
}
