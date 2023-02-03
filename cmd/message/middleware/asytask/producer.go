package asytask

import (
	"context"
	"encoding/json"
	"log"
	// "runedance_douyin/cmd/message/dal/db_redis"
	// "time"

	"github.com/hibiken/asynq"
)

// add new sync tack into the queue
func AddNewTask(ctx context.Context, userId int64, toUserId int64, delaySec int) error{
	m := TaskPlayload{
		UserId: userId,
		ToUserId: toUserId,
	}

	jsonStr, err := json.Marshal(m)
	if(err != nil){
		panic(err)
	}
	print(jsonStr)
	newtask := asynq.NewTask("sync", jsonStr)
	if(AsyClient == nil){
		log.Printf(string(newtask.Payload()))
	}
	
	// taskInfo, err := AsyClient.Enqueue(newtask, asynq.ProcessIn(time.Duration(delaySec)))
	
	// if(err != nil){
	// 	log.Fatalf(string(taskInfo.Payload))
	// 	// log.Printf(string(taskInfo.Payload))
	// 	return err
	// }
	return err
	// // deal with repeated task 
	// inspector := asynq.NewInspector(RdbConn)

	// // get all pending repeated task
	// pendingTaskList, _ := db_redis.GetPendingTaskIDs(ctx, userId, toUserId)
	// qname, ok := asynq.GetQueueName(ctx)
	// if(ok && pendingTaskList != nil){
	// 	for _, val:= range pendingTaskList {
	// 		inspector.DeleteTask(qname, val)			// delete pending sync task
	// 	}
	// 	db_redis.ClearTaskList(ctx, m.UserId, m.ToUserId)
	// }
	// if err := db_redis.AddNewTask(ctx, m.UserId, m.ToUserId, taskInfo.ID); err != nil{
	// 	return err
	// }
	// return nil
}