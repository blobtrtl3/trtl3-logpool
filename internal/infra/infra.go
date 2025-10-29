package infra

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedistClient(ctx context.Context) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

