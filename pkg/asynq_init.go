package pkg

import (
	"github.com/hibiken/asynq"

)


var srv *asynq.Server
var asyClient *asynq.Client
var inspector *asynq.Inspector
var asynqConn asynq.RedisClientOpt

func InitAsynq() {
	asynqConn = asynq.RedisClientOpt{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}
	
	srv = asynq.NewServer(asynqConn, asynq.Config{
		Concurrency: 10,
	})
}

func GetAsynqServer() *asynq.Server{
	return srv
}

func GetAsynqClient() *asynq.Client{
	asyClient = asynq.NewClient(asynqConn)
	return asyClient
}

func GetAsynqInspector() *asynq.Inspector {
	inspector = asynq.NewInspector(asynqConn)
	return inspector
}