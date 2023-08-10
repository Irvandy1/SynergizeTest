package config

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis"
)

type Config struct {
	DB          *sql.DB
	RedisClient *redis.Client
}

func (cfg Config) init() (err error) {
	cfg.DB, err = initPostgres()
	if err != nil {
		fmt.Println(err)
	}

	cfg.RedisClient = initRedis()

	return
}
