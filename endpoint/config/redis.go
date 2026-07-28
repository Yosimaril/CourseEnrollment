package config

import (
	"context"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var (
	Redis *redis.Client
	Ctx   = context.Background()
)

func ConnectRedis() {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}

	addr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")

	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: os.Getenv("REDIS_PASS"),
		DB:       db,
	})

	if err := Redis.Ping(Ctx).Err(); err != nil {
		panic(err)
	}
}
