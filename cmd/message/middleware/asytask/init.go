package asytask

import (
	"context"
	"encoding/json"
	"runedance_douyin/cmd/message/dal/db_redis"
	"github.com/hibiken/asynq"
	"time"

	"log"
)

var RdbConn asynq.RedisClientOpt
var asyClient *asynq.Client

func Init() {
	RdbConn = asynq.RedisClientOpt{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}

	asyClient = asynq.NewClient(RdbConn)

	srv := asynq.NewServer(RdbConn, asynq.Config{
		Concurrency: 10,
	})
	if err := srv.Run(asynq.HandlerFunc(SyncTaskHandler)); err != nil {
		log.Fatal(err)
	}
}

func NewSyncTask(userId int64, toUserId int64) *asynq.Task {
	m := make(map[string]int64)
	m["userId"] = userId
	m["toUserId"] = toUserId
	jsonStr, err := json.Marshal(m)
	if(err != nil){
		panic(err)
	}
	return asynq.NewTask("sync", jsonStr)
}

// method to handle task
func SyncTaskHandler(ctx context.Context, task *asynq.Task) error {
	m := make(map[string]int64)
	err := json.Unmarshal(task.Payload(), &m)
	if(err != nil){
		return err
	}
	return TransMsgFromRedisToDB(ctx, m["userId"], m["toUserId"])	
}


// add new sync tack into the queue
func AddNewTask(ctx context.Context, task *asynq.Task, delaySec int) error{
	taskInfo, err := asyClient.Enqueue(task, asynq.ProcessIn(time.Duration(delaySec)))
	if(err != nil){
		return err
	}
	
	// deal with repeated task 
	inspector := asynq.NewInspector(RdbConn)
	m := make(map[string]int64)
	if err := json.Unmarshal(task.Payload(), &m); err != nil {
		return err
	}
	// get all pending repeated task
	pendingTaskList, _ := db_redis.GetPendingTaskIDs(ctx, m["userId"], m["toUserId"])
	qname, ok := asynq.GetQueueName(ctx)
	if(ok && pendingTaskList != nil){
		for _, val:= range pendingTaskList {
			inspector.DeleteTask(qname, val)			// delete pending sync task
		}
		db_redis.ClearTaskList(ctx, m["userId"], m["toUserId"])
	}
	if err := db_redis.AddNewTask(ctx, m["userId"], m["toUserId"], taskInfo.ID); err != nil{
		return err
	}
	return nil
}