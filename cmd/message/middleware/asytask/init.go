package asytask

import (
	"github.com/hibiken/asynq"
	"log"
	"github.com/redis/go-redis/v9"
)


var Rdb *redis.Client

func InitServer() {
	RdbConn := asynq.RedisClientOpt{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}
	
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	srv := asynq.NewServer(RdbConn, asynq.Config{
		Concurrency: 10,
	})
	mux := asynq.NewServeMux()
	mux.HandleFunc("sync", SyncTaskHandler)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

type TaskPlayload struct {
	UserId  int64
	ToUserId int64
	CreateTime string
}




