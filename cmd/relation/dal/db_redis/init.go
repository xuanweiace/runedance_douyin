package db_redis

import (
	constants "runedance_douyin/pkg/consts"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func Init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
