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

func getRdb() *redis.Client{
	Init()
	return Rdb
}


var RdbCluster *redis.ClusterClient

func InitCluster() {
	RdbCluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:     []string {"localhost:6380", "localhost:6381", "localhost:6382", "localhost:6383", "localhost:6384", "localhost:6385"},
		// Password: "", // no password set
		// DB:       0,  // use default DB
	})
}

func getRdbCluster() *redis.ClusterClient{
	InitCluster()
	return RdbCluster
}