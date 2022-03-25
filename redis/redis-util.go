package redis

import (
	"context"
	"fmt"
)

func SetKey(key, val string) {

	client := GetRedis()
	err := client.Set(context.Background(), key, val, 0).Err()
	if err != nil {
		fmt.Println("Error set redis : ", err.Error())
	}
}

func GetKey(key string) string {

	client := GetRedis()
	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		fmt.Println("Error set redis : ", err.Error())
		return ""
	}
	return val
}

func DelKey(key string) {
	client := GetRedis()
	client.Del(context.Background(), key)
}
