package db

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func RedisConn() (*redis.Client,error) {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
		DB: 0,
		DialTimeout: 100 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
	})

	if _,err := client.Ping(context.Background()).Result(); err != nil {
		return nil,err
	}

	return client,nil
}

