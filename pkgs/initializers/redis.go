package initializers

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func ConnectToRedis() (*redis.Client, error) {
	ctx := context.Background()
	fmt.Println("Testing Golang Redis")

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis", err)
		return nil, err
	} else {
		fmt.Println("Connected to Redis", pong)
		return client, nil
	}
}
