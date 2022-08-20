package connection

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func ConnectToRedis() *redis.Client {
	// inisialisasi koneksi
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //hostname
		Password: "",               //password
		DB:       0,
	})

	ctx := context.Background()
	result, err := redis.Ping(ctx).Result()
	if err != nil || result != "PONG" {
		log.Println("not connect error =>", err)
		log.Println("Redis not connected !!")
	}
	log.Println("Redis Connected")
	return redis
}
