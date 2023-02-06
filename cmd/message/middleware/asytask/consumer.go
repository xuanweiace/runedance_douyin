package asytask

import (
	"context"
	"encoding/json"

	"strconv"

	"log"
	"runedance_douyin/cmd/message/dal/db_mysql"
	"runedance_douyin/cmd/message/dal/db_redis"

	"runedance_douyin/kitex_gen/message"
	"runedance_douyin/pkg/tools"

	"github.com/hibiken/asynq"
)

// method to handle task
func SyncTaskHandler(ctx context.Context, task *asynq.Task) error {
	var m TaskPlayload
	err := json.Unmarshal(task.Payload(), &m)
	log.Printf(string(task.Payload()))
	if(err != nil){
		return err
	}
	log.Printf("task created at: " + m.CreateTime)
	// return TransMsgFromRedisToDB(ctx, m.UserId, m.ToUserId)
	// log.Printf(strconv.FormatInt(m.UserId, 10))
	// log.Printf(strconv.FormatInt(m.ToUserId, 10))	
	
	err2 := TransMsgFromRedisToDB(ctx, m.UserId, m.ToUserId)
	return err2
}



// sync DB with redis
func TransMsgFromRedisToDB(ctx context.Context, userId int64, toUserId int64) error{
	// log.Printf(strconv.FormatInt(userId, 10))

	// start new goroutine, need to init redis and mysql
	db_mysql.Init()
	db_redis.Init()
	
	redis_msg, err:= db_redis.GetMessageChatJson(ctx, userId, toUserId)
	if(err != nil){
		return err 
	}
	if(redis_msg != nil){
		log.Printf("getting message string from redis....")
	}
	
	var messageRecordList []*db_mysql.MessageRecord

	if(len(redis_msg) == 0){								
		return nil
	}

	// // get latest sync timestamp 
	latestTime, err := db_redis.GetTimestampOfLatestMysql(ctx, userId, toUserId)
	if(err != nil){
		return err
	}
	
	var newLatest int64
	for _, val := range redis_msg {
		var msg message.Message
		err := json.Unmarshal([]byte(val), &msg)			// decode json into Message struct
		if(err != nil){
			continue
		}
		if(msg.Id < latestTime){
			continue
		}
		msgRecord := db_mysql.MessageRecord{
			UserToUser: tools.GenerateKeyname(userId, toUserId),
			Content : msg.Content,
			CreateTime : msg.CreateTime, 
		}

		messageRecordList = append(messageRecordList, &msgRecord)
		newLatest = msg.Id
	}
	// // update latest sync timestamp 
	db_redis.SetTimestampOfLatestMysql(ctx, userId, toUserId, newLatest)
	log.Printf("new latest time: " + strconv.FormatInt(newLatest, 10))
	return db_mysql.InsertMessage(ctx, messageRecordList)
	// return nil
}