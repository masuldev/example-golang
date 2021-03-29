package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var (
	ErrNotFoundEnv = errors.New("You must make .env file")
)

func ConnectionRedis() *redis.Client {
	addr, pass := makeRedisSource()

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	return rdb
}

func makeRedisSource() (string, string) {
	pass := os.Getenv("REDIS_PASSWORD")
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	return fmt.Sprintf("%s:%s", host, port), pass
}
