package connection

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func ConnectRedis() *redis.Client {
	// Inisialisasi Redis Connection
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //hostname
		Password: "",               //password
		DB:       0,
	})

	ctx := context.Background()
	// test connection redis
	msg, err := redis.Ping(ctx).Result()

	// checked nil or return PONG from Ping(ctx)
	// return PONG if no argument is provided
	if err != nil || msg != "PONG" {
		log.Println("Error =>", err)
		log.Println("Redis Not Connect")
	} else {
		log.Println("Redis Connected")
	}
	return redis
}
