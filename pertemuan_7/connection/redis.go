package connection

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func ConnectToRedis() *redis.Client {
	// insialisasi Koneksi
	redis := redis.NewClient(&redis.Options{
		Addr:     "ec2-3-217-134-177.compute-1.amazonaws.com:17599",                   //hostname
		Password: "p0a639e9ccf0dafca1c8d23cc40cd6a34297712ddfd90f9ab59ab11e81f97c959", //password
		DB:       0,                                                                   //bisa 0,1,2,3,4,5,6,7,8,9 dll
	})

	ctx := context.Background()
	msg, err := redis.Ping(ctx).Result() // Test koneksi Redis nya (nyala atau engga)

	if err != nil || msg != "PONG" {
		log.Println("not conect error =>", err)
		log.Println("Redis Not Connected !!")
	} else {
		log.Println("Redis Connected")
	}
	return redis
}
