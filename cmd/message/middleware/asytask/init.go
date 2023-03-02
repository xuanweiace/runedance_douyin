package asytask

import (
	"github.com/hibiken/asynq"
	"log"
	"runedance_douyin/pkg"
)


func InitServer() {
	pkg.InitAsynq()
	srv := pkg.GetAsynqServer()
	mux := asynq.NewServeMux()
	mux.HandleFunc("sync", SyncTaskHandler)
	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

func getClient() *asynq.Client{
	return pkg.GetAsynqClient()
}

func getInspector() *asynq.Inspector {
	return pkg.GetAsynqInspector()
}

type TaskPlayload struct {
	UserId  int64
	ToUserId int64
	CreateTime string
}




