package asytask

import (
	"context"
	"encoding/json"
	"hash/fnv"
	"log"
	"runedance_douyin/cmd/message/dal/db_redis"
	"time"
	"strconv"
	"github.com/hibiken/asynq"
	"runedance_douyin/pkg/tools"
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
	keyname := tools.GenerateKeyname(userId, toUserId)
	qname := getHashQueue(keyname)

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
	log.Printf("the current queue is: " + qname)
	log.Printf("the current task is: " + taskInfo.ID)

	// deal with repeated task 
	// // get all pending repeated task
	pendingTaskList, err := db_redis.GetPendingTaskIDs(ctx, userId, toUserId)
	if(err != nil){
		return err
	}

	inspector := getInspector()
	defer inspector.Close()
	if(pendingTaskList != nil){
		for _, val:= range pendingTaskList {
			inspector.DeleteTask(qname, val)					// delete pending sync task
			log.Printf("delete repeated task: " + val)
		}
		db_redis.ClearTaskList(ctx, m.UserId, m.ToUserId)
	}

	if err := db_redis.AddNewTask(ctx, m.UserId, m.ToUserId, taskInfo.ID); err != nil{
		return err
	}
	return nil
}

func getHashQueue(key string) string {	
	hash := fnv.New64()
	hash.Write([]byte(key))
	seq := hash.Sum64() % 29
	return "queue" + strconv.FormatUint(seq , 10) 
}