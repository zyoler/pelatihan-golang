package helperRedis

import (
	"context"
	r "randi_firmansyah/connections/redis"

	"github.com/go-redis/redis/v8"
)

var (
	REDIS *redis.Client
	ctx   = context.Background()
)

func init() {
	REDIS = r.ConnectToRedis()
}

func ClearRedis(key string) {
	// del redis
	REDIS.Del(ctx, key)
}

func SetRedis(keyRedis string, result []byte) error {
	// set ke redis
	if err := REDIS.Set(ctx, keyRedis, (result), 0).Err(); err != nil {
		return err
	}
	return nil
}

func GetRedis(key string) (string, error) {
	// get redis data by key
	response, err := REDIS.Get(ctx, key).Result()
	return response, err
}
