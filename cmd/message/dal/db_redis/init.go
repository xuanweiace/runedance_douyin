package db_redis

import (
    "github.com/redis/go-redis/v9"
)

// var ctx = context.Background()

var Rdb *redis.Client

func Init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}