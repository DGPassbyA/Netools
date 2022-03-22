package database

import (
	"context"
	"log"
	"main/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func RedisInit() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPwd, // password
		DB:       0,               // use default DB
		PoolSize: 100,             // pool size
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Init Redis successfully.")
}
