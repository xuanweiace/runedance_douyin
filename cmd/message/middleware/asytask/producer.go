package asytask

import (
	"context"
	"encoding/json"
	"log"
	"runedance_douyin/cmd/message/dal/db_redis"
	"time"
	"github.com/hibiken/asynq"

)

// add new sync tack into the queue
func AddNewTask(ctx context.Context, userId int64, toUserId int64) error{
	curTime := time.Now()
	m := TaskPlayload{
		UserId: userId,
		ToUserId: toUserId,
		CreateTime: curTime.String(),					
	}

	jsonStr, err := json.Marshal(m)
	log.Print(string(jsonStr))

	if(err != nil){
		panic(err)
	}

	newtask := asynq.NewTask("sync", jsonStr)

	// create asynqClient to push task
	asynqClient := getClient()
	defer asynqClient.Close()	
	taskInfo, err := asynqClient.Enqueue(newtask, asynq.ProcessAt(curTime.Add(time.Minute)), asynq.MaxRetry(1))
	if(err != nil){
		log.Fatalf(string(taskInfo.Payload))
		return err
	}
	log.Printf("add task to queue")
	log.Printf("the current task is: " + taskInfo.ID)

	// deal with repeated task 
	// // get all pending repeated task
	db_redis.InitCluster()
	pendingTaskList, err := db_redis.GetPendingTaskIDs(ctx, userId, toUserId)
	if(err != nil){
		return err
	}

	inspector := getInspector()
	defer inspector.Close()
	if(pendingTaskList != nil){
		for _, val:= range pendingTaskList {
			inspector.DeleteTask(taskInfo.Queue, val)					// delete pending sync task
			log.Printf("delete repeated task: " + val)
		}
		db_redis.ClearTaskList(ctx, m.UserId, m.ToUserId)
	}

	if err := db_redis.AddNewTask(ctx, m.UserId, m.ToUserId, taskInfo.ID); err != nil{
		return err
	}
	return nil
}
