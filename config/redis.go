package config

import (
	"os"

	"github.com/go-redis/redis"
)

func initRedis() (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("localhost:6379"),
		Password: os.Getenv(""),
		DB:       0,
	})
	return
}
