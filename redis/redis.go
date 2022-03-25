package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func init() {
	client = NewRedisConfig()
}

func NewRedisConfig() *redis.Client {

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	client := redis.NewClient(
		&redis.Options{
			Addr:     redisHost + ":" + redisPort,
			Password: "",
			DB:       0,
		},
	)
	_, err := client.Ping(context.Background()).Result()
	if err == nil {
		fmt.Println("Redis connected !")
	} else {
		fmt.Println("Ping redis - err : ", err)
	}
	return client

}

func GetRedis() *redis.Client {
	if client == nil {
		client = NewRedisConfig()
		return client
	}
	return client
}
