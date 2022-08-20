package connection

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func ConnectToRedis() *redis.Client {
	// insialisasi Koneksi
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //hostname
		Password: "",               //password
		DB:       0,                //bisa 0,1,2,3,4,5,6,7,8,9 dll
	})

	ctx := context.Background()
	msg, err := redis.Ping(ctx).Result()

	if err != nil || msg != "PONG" {
		log.Println("not conect error =>", err)
		log.Println("Redis Not Connected !!")
	} else {
		log.Println("Redis Connected")
	}
	return redis
}
