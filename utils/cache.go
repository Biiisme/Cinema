package utils

import (
	"log"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var Cache *redis.Client

func InitCache() error {
	Cache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := Cache.Ping().Result(); err != nil {
		log.Fatal("Redis: không có kết nối", zap.Error(err))
		return err
	} else {
		return nil
	}
}
