package config

import (
	"os"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func initRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("localhost:6379"),
		Password: os.Getenv(""),
		DB:       0,
	})
	return
}
