package asytask

import (
	"github.com/hibiken/asynq"
	"log"
)

var RdbConn asynq.RedisClientOpt
var AsyClient *asynq.Client

func Init() {
	RdbConn = asynq.RedisClientOpt{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}
	AsyClient = asynq.NewClient(RdbConn)
	srv := asynq.NewServer(RdbConn, asynq.Config{
		Concurrency: 10,
	})
	if err := srv.Run(asynq.HandlerFunc(SyncTaskHandler)); err != nil {
		log.Fatal(err)
	}
}

type TaskPlayload struct {
	UserId  int64
	ToUserId int64
}




