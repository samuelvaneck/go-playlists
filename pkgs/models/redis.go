package models

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisClient(r *redis.Client) RedisClient {
	ctx := context.Background()
	return RedisClient{Client: r, Ctx: ctx}
}

func (r *RedisClient) SetInRedis(key, val string) {
	err := r.Client.Set(r.Ctx, key, val, time.Duration(60*time.Minute)).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (r *RedisClient) GetFromRedis(key string) (string, error) {
	val, err := r.Client.Get(r.Ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return val, nil
}

func (r *RedisClient) GetAllKeys(ctx context.Context, key string) []string {
	keys := []string{}

	iter := r.Client.Scan(ctx, 0, key, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	return keys
}
