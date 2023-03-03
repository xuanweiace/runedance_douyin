package db_redis

import (
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func Init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
