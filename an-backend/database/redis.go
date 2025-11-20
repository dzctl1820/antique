package database

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// 测试Redis连接
	ctx := context.Background()
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Printf("Redis连接失败: %v", err)
	} else {
		log.Printf("Redis连接成功: %s", pong)
	}
}

func InsertRedis(key string, value string) error {
	ctx := context.Background()
	expiration := 5 * time.Minute
	err := redisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	log.Printf("验证码存储成功: key=%s, value=%s", key, value)
	return nil
}

func FindRedis(key string) (string, error) {
	result, err := redisClient.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
