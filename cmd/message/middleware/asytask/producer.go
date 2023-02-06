package asytask

import (
	"context"
	"encoding/json"
	"log"
	"runedance_douyin/cmd/message/dal/db_redis"
	"time"
	"github.com/hibiken/asynq"
)
var AsyClient *asynq.Client
// add new sync tack into the queue
func AddNewTask(ctx context.Context, userId int64, toUserId int64) error{

	RdbConn := asynq.RedisClientOpt{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}
	

	AsyClient = asynq.NewClient(RdbConn)
	defer AsyClient.Close()

	m := TaskPlayload{
		UserId: userId,
		ToUserId: toUserId,
		CreateTime: time.Now().String(),
	}

	jsonStr, err := json.Marshal(m)
	log.Printf(string(jsonStr))

	if(err != nil){
		panic(err)
	}
	newtask := asynq.NewTask("sync", jsonStr)
	
	taskInfo, err := AsyClient.Enqueue(newtask, asynq.ProcessAt(time.Now().Add(time.Minute)), asynq.MaxRetry(1))
	if(err != nil){
		log.Fatalf(string(taskInfo.Payload))
		return err
	}
	log.Printf("add task to queue")
	qname := taskInfo.Queue
	log.Printf("the current queue is: " + qname)

	// deal with repeated task 
	inspector := asynq.NewInspector(RdbConn)

	// // get all pending repeated task
	pendingTaskList, err := db_redis.GetPendingTaskIDs(ctx, userId, toUserId)
	if(err != nil){
		return err
	}

	if(pendingTaskList != nil){
		for _, val:= range pendingTaskList {
			inspector.DeleteTask(qname, val)			// delete pending sync task
			log.Printf("delete task:" + val)
		}
		db_redis.ClearTaskList(ctx, m.UserId, m.ToUserId)
	}

	if err := db_redis.AddNewTask(ctx, m.UserId, m.ToUserId, taskInfo.ID); err != nil{
		return err
	}
	return nil
}